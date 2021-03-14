package service

import (
	"context"

	"github.com/jmoiron/sqlx"
	"gitlab.com/cadaverine/pim-service/models"
)

// SaveCatalog ...
func (s *PimService) SaveCatalog(ctx context.Context, cg *models.Catalog) error {
	for _, shop := range cg.Shops {
		err := s.db.InTx(ctx, nil, func(tx *sqlx.Tx) error {
			shopID, err := s.SaveShop(ctx, shop)
			if err != nil {
				return err
			}

			err = s.SaveCategories(ctx, tx, shopID, shop.Categories.Categories)
			if err != nil {
				return err
			}

			err = s.SaveProducts(ctx, tx, shopID, shop.Offers.Offers)
			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}
