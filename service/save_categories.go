package service

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"gitlab.com/cadaverine/pim-service/models"
)

// SaveCategories ...
func (s *PimService) SaveCategories(ctx context.Context, tx *sqlx.Tx, shopID int, categories []models.Category) error {
	return s.db.InTx(ctx, tx, func(t *sqlx.Tx) error {
		err := s.deleteCategories(ctx, tx, shopID)
		if err != nil {
			return errors.Wrap(err, "failed to delete categories")
		}

		err = s.addCategories(ctx, tx, shopID, categories)
		if err != nil {
			return errors.Wrapf(err, "failed to add categories for shopID '%v'", shopID)
		}

		return nil
	})
}

func (s *PimService) addCategories(ctx context.Context, tx *sqlx.Tx, shopID int, categories []models.Category) error {
	for _, category := range categories {
		err := s.addCategory(ctx, tx, shopID, category)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *PimService) addCategory(ctx context.Context, tx *sqlx.Tx, shopID int, category models.Category) error {
	const query = `
		insert into product_information.categories (shop_id, item_id, parent_id, name)
		values ($1, $2, $3, $4) on conflict (shop_id, item_id) do
		update set
			shop_id = excluded.shop_id,
			item_id = excluded.item_id,
			parent_id = excluded.parent_id,
			name = excluded.name
	`

	return s.db.InTx(ctx, tx, func(t *sqlx.Tx) error {
		_, err := tx.Exec(query,
			shopID,
			category.ID,
			category.ParentID,
			category.Title,
		)
		return err
	})
}

func (s *PimService) deleteCategories(ctx context.Context, tx *sqlx.Tx, shopID int) error {
	const query = `
		delete from product_information.categories
		where shop_id = $1
	`

	return s.db.InTx(ctx, tx, func(t *sqlx.Tx) error {
		_, err := tx.Exec(query, shopID)
		return err
	})
}
