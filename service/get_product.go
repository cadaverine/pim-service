package service

import (
	"context"

	"github.com/jmoiron/sqlx"
	gen "gitlab.com/cadaverine/pim-service/gen/pim-service"
	"gitlab.com/cadaverine/pim-service/models"
)

func (s *PimService) GetProduct(ctx context.Context, req *gen.ProductIDs) (*gen.Product, error) {
	id, shopID := req.GetID(), req.GetShopID()

	var err error
	var product *models.Offer

	err = s.db.InTx(ctx, nil, func(tx *sqlx.Tx) error {
		product, err = getProduct(ctx, tx, id, shopID)
		if err != nil {
			return err
		}

		offersMap, err := s.getOffersParamsByIDs(ctx, tx, []int{product.ID})
		if err != nil {
			return err
		}

		product.Param = offersMap[product.ID]

		return nil
	})
	if err != nil {
		return nil, err
	}

	return repackOfferToProto(*product), nil
}

func getProduct(ctx context.Context, tx *sqlx.Tx, id string, shopID int32) (*models.Offer, error) {
	const query = `
		select
			product.id, product.item_id, product.shop_id, product.name,
			product.available, product.type, product.url, product.price,
			product.currency_code, product.vendor, product.description
		from product_information.products product
		where item_id = $1 and shop_id = $2
		limit 1;
	`

	var offer models.Offer

	err := tx.QueryRow(query, id, shopID).Scan(
		&offer.ID,
		&offer.ItemID,
		&offer.ShopID,
		&offer.Name,
		&offer.Available,
		&offer.Type,
		&offer.URL,
		&offer.Price,
		&offer.CurrencyID,
		&offer.Vendor,
		&offer.Description,
	)
	if err != nil {
		return nil, err
	}

	return &offer, nil
}
