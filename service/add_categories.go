package service

import (
	"context"

	"github.com/jackc/pgx/v4"
)

// AddCategories ...
func (s *PimService) CreateCategories(ctx context.Context, tx pgx.Tx, sellerID int, categories []category) error {
	err := s.db.InTx(ctx, tx, func(tx pgx.Tx) error {
		for _, category := range categories {
			err := s.createCategory(ctx, tx, sellerID, category)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *PimService) createCategory(ctx context.Context, tx pgx.Tx, sellerID int, cg category) error {

	const query = `insert into products.`

	_, err := tx.Exec(ctx, query, sellerID, cg.ID, cg.ParentID, cg.Title)
	if err != nil {
		return err
	}

	return nil
}
