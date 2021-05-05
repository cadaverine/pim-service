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

// CreateProduct добавить продукт для магазина
func (s *PimService) CreateProduct(ctx context.Context, req *gen.Product) (*empty.Empty, error) {
	product := repackProductFromProto(req)
	if product == nil {
		return nil, errors.New("fields requiered")
	}

	return nil, s.db.InTx(ctx, nil, func(tx *sqlx.Tx) error {
		id, err := createProduct(ctx, tx, product.ShopID, *product)
		if err != nil {
			return err
		}

		return s.addProductAttributes(ctx, tx, id, product.Param)
	})
}

func createProduct(ctx context.Context, tx *sqlx.Tx, shopID int, offer models.Offer) (int, error) {
	const query = `
		insert into product_information.products (
			item_id,
			shop_id,
			name,
			available,
			type,
			url,
			price,
			currency_code,
			vendor,
			description
		) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
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
		return 0, errors.Wrapf(err, "failed to save product: $+v", offer)
	}

	return productID, err
}

func repackProductFromProto(src *gen.Product) *models.Offer {
	if src == nil {
		return nil
	}

	return &models.Offer{
		ShopID:      int(src.ShopID),
		ItemID:      src.ItemID,
		Available:   src.Available,
		Type:        src.Type,
		URL:         src.Url,
		Price:       string(src.Price),
		CurrencyID:  src.CurrencyID,
		CategoryID:  src.CategoryID,
		Name:        src.Name,
		Vendor:      src.Vendor,
		Description: src.Description,
		Param:       repackParamsFromProto(src.Params),
	}
}

func repackParamsFromProto(src []*gen.Product_Param) []models.Param {
	res := make([]models.Param, 0, len(src))

	for _, item := range src {
		if item == nil {
			continue
		}

		value, _ := item.GetValue().MarshalJSON()

		res = append(res, models.Param{
			Name:  item.Name,
			Value: string(value),
			Type:  item.Type,
		})
	}

	return res
}
