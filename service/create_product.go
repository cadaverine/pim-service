package service

import (
	"context"
	"errors"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jmoiron/sqlx"
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
		id, err := s.addProduct(ctx, tx, product.ShopID, *product)
		if err != nil {
			return err
		}

		return s.addProductAttributes(ctx, tx, id, product.Param)
	})
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

		res = append(res, models.Param{
			Name:  item.Name,
			Value: item.GetValue().String(),
			Type:  item.Type,
		})
	}

	return res
}
