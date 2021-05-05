package service

import (
	"context"
	"strconv"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	gen "gitlab.com/cadaverine/pim-service/gen/pim-service"
	"gitlab.com/cadaverine/pim-service/models"
)

func (s *PimService) UpdateProduct(ctx context.Context, req *gen.Product) (*empty.Empty, error) {
	product := repackProductFromProto(req)
	if product == nil {
		return nil, errors.New("fields requiered")
	}

	return nil, s.db.InTx(ctx, nil, func(tx *sqlx.Tx) error {
		id, err := updateProduct(ctx, tx, product.ShopID, *product)
		if err != nil {
			return err
		}

		err = deleteProductAttributes(ctx, tx, id)
		if err != nil {
			return err
		}

		return s.addProductAttributes(ctx, tx, id, product.Param)
	})
}

func updateProduct(ctx context.Context, tx *sqlx.Tx, shopID int, offer models.Offer) (int, error) {
	const query = `
		update product_information.products
		set
			name = $3,
			available = $4,
			type = $5,
			url = $6,
			price = $7,
			currency_code = $8,
			vendor = $9,
			description = $10
		where item_id = $1 and shop_id = $2
		returning id;
	`

	var productID int

	price, _ := strconv.Atoi(offer.Price)

	err := tx.Get(&productID, query,
		offer.ItemID,
		shopID,
		offer.Name,
		offer.Available,
		offer.Type,
		offer.URL,
		price,
		offer.CurrencyID,
		offer.Vendor,
		offer.Description,
	)
	if err != nil {
		return 0, errors.Wrapf(err, "failed to update product: $+v", offer)
	}

	return productID, err
}

func deleteProductAttributes(ctx context.Context, tx *sqlx.Tx, id int) error {
	const query = `
		delete from product_information.products_attributes
		where product_id = $1
	`

	_, err := tx.Exec(query, id)
	return err
}
