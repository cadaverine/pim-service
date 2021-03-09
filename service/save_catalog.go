package service

import (
	"context"

	"github.com/jackc/pgx/v4"
)

func (s *PimService) SaveCatalog(ctx context.Context, cg *catalog) error {
	return s.db.InTx(ctx, nil, func(tx pgx.Tx) error {
		for _, shop := range cg.Shops {
			err := s.UpdateCategories(ctx, tx, shop.Name, shop.Categories.Categories)
			if err != nil {
				return err
			}

			err = s.UpdateProducts(ctx, tx, shop.Name, shop.Offers.Offers)
			if err != nil {
				return err
			}
		}

		return nil
	})
}
