package service

import (
	"context"
	"errors"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jmoiron/sqlx"
	gen "gitlab.com/cadaverine/pim-service/gen/pim-service"
	"gitlab.com/cadaverine/pim-service/models"
)

func (s *PimService) CreateCategory(ctx context.Context, req *gen.Category) (*empty.Empty, error) {
	id := req.GetID()
	name := req.GetName()
	shopID := req.GetShopID()
	parentID := req.GetParentID()

	if name == "" || id == 0 || shopID == 0 {
		return &empty.Empty{}, errors.New("request validation error")
	}

	return &empty.Empty{}, s.db.InTx(ctx, nil, func(tx *sqlx.Tx) error {
		return createCategory(ctx, tx, shopID, models.Category{
			ID:       int(id),
			ParentID: int(parentID),
			Title:    name,
		})
	})
}

func createCategory(ctx context.Context, tx *sqlx.Tx, shopID int32, category models.Category) error {
	const query = `
		insert into product_information.categories (shop_id, item_id, parent_id, name)
		values ($1, $2, $3, $4)
	`

	_, err := tx.Exec(query, shopID, category.ID, category.ParentID, category.Title)
	return err
}
