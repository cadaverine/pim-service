package service

import (
	"context"

	"github.com/jmoiron/sqlx"
	gen "gitlab.com/cadaverine/pim-service/gen/pim-service"
	"gitlab.com/cadaverine/pim-service/models"
)

func (s *PimService) GetCategory(ctx context.Context, req *gen.IDs) (*gen.Category, error) {
	id, shopID := req.GetID(), req.GetShopID()

	var res gen.Category

	return &res, s.db.InTx(ctx, nil, func(tx *sqlx.Tx) error {
		category, err := getCategory(ctx, tx, id, shopID)
		if err != nil {
			return err
		}

		res.ID = int32(category.ID)
		res.ParentID = int32(category.ParentID)
		res.Name = category.Title

		return nil
	})
}

func getCategory(ctx context.Context, tx *sqlx.Tx, id, shopID int32) (*models.Category, error) {
	const query = `
		select item_id, parent_id, name
		from product_information.categories
		where item_id = $1 and shop_id = $2
		limit 1;
	`

	var category models.Category

	err := tx.QueryRow(query, id, shopID).Scan(
		&category.ID,
		&category.ParentID,
		&category.Title,
	)
	if err != nil {
		return nil, err
	}

	return &category, nil
}
