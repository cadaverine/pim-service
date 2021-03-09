package service

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

// UpdateCategories updates categories
func (s *PimService) UpdateCategories(ctx context.Context, tx pgx.Tx, shop string, categories []category) error {
	return s.db.InTx(ctx, tx, func(t pgx.Tx) error {
		err := s.deleteCategories(ctx, tx)
		if err != nil {
			return errors.Wrap(err, "failed to delete categories")
		}

		err = s.addCategories(ctx, tx, shop, categories)
		if err != nil {
			return errors.Wrapf(err, "failed to add categories for shop '%s'", shop)
		}

		return nil
	})
}

func (s *PimService) addCategories(ctx context.Context, tx pgx.Tx, shop string, categories []category) error {
	const query = `
		insert into products.categories (id, parent_id, name)
		values ($1, $2, $3) on conflict (id) do
		update set
			id = excluded.id,
			parent_id = excluded.parent_id,
			name = excluded.name,
			deleted_at = null;
	`

	batch := &pgx.Batch{}

	for _, category := range categories {
		batch.Queue(query, category.ID, category.ParentID, category.Title)
	}

	return s.db.InTx(ctx, tx, func(t pgx.Tx) error {
		br := tx.SendBatch(ctx, batch)
		defer br.Close()

		for i := 0; i < batch.Len(); i++ {
			_, err := br.Exec()
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (s *PimService) deleteCategories(ctx context.Context, tx pgx.Tx) error {
	const query = `
		update products.categories
		set deleted_at = now();
	`

	return s.db.InTx(ctx, tx, func(t pgx.Tx) error {
		_, err := tx.Exec(ctx, query)
		return err
	})
}
