package service

import (
	"context"
	"errors"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jmoiron/sqlx"
	gen "gitlab.com/cadaverine/pim-service/gen/pim-service"
)

func (s *PimService) DeleteCategory(ctx context.Context, req *gen.IDs) (*empty.Empty, error) {
	id, shopID := req.GetID(), req.GetShopID()

	if id == 0 || shopID == 0 {
		return &empty.Empty{}, errors.New("request validation error")
	}

	return &empty.Empty{}, s.db.InTx(ctx, nil, func(tx *sqlx.Tx) error {
		return deleteCategory(ctx, tx, id, shopID)
	})
}

func deleteCategory(ctx context.Context, tx *sqlx.Tx, id, shopID int32) error {
	const query = `
		delete from product_information.categories
		where item_id = $1 and shop_id = $2
	`

	_, err := tx.Exec(query, id, shopID)
	return err
}
