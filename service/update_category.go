package service

import (
	"context"
	"errors"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jmoiron/sqlx"
	gen "gitlab.com/cadaverine/pim-service/gen/pim-service"
)

func (s *PimService) UpdateCategory(ctx context.Context, req *gen.Category) (*empty.Empty, error) {
	id := req.GetID()
	name := req.GetName()
	shopID := req.GetShopID()
	parentID := req.GetParentID()

	if name == "" || id == 0 || shopID == 0 {
		return &empty.Empty{}, errors.New("request validation error")
	}

	return &empty.Empty{}, s.db.InTx(ctx, nil, func(tx *sqlx.Tx) error {
		return s.updateCategory(ctx, tx, id, shopID, parentID, name)
	})
}

func (s *PimService) updateCategory(ctx context.Context, tx *sqlx.Tx, id, shopID, parentID int32, name string) error {
	const query = `
		update product_information.categories
		set name = $1, parent_id = $2
		where item_id = $3 and shop_id = $4
	`

	_, err := tx.Exec(query, name, parentID, id, shopID)
	return err
}
