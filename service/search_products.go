package service

import (
	"context"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	gen "gitlab.com/cadaverine/pim-service/gen/go/api/pim-service"
	"gitlab.com/cadaverine/pim-service/models"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
)

const (
	defaultLimit = 15
	maxLimit     = 200
)

func (s *PimService) SearchProducts(ctx context.Context, req *gen.SearchRequest) (*gen.Products, error) {
	shopID := int(req.GetShopID())
	searchTerm := req.GetSearchTerm()

	limit, offset := int(req.GetMeta().GetLimit()), int(req.GetMeta().GetOffset())

	if limit == 0 || limit > maxLimit {
		limit = defaultLimit
	}

	var err error
	var offers []models.Offer

	err = s.db.InTx(ctx, nil, func(tx *sqlx.Tx) error {
		offers, err = s.searchOffers(ctx, tx, shopID, searchTerm, limit, offset)
		if err != nil {
			return err
		}

		offersIDs := make([]int, len(offers))
		for i := range offers {
			offersIDs[i] = offers[i].ID
		}

		offersMap, err := s.getOffersParamsByIDs(ctx, tx, offersIDs)
		if err != nil {
			return err
		}

		for i := range offers {
			offers[i].Param = offersMap[offers[i].ID]
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &gen.Products{
		Products: repackOffersToProto(offers),
		Meta: &gen.ScrollDescriptor{
			Limit:  int32(limit),
			Offset: int32(offset),
		},
	}, nil
}

func repackOffersToProto(src []models.Offer) []*gen.Product {
	res := make([]*gen.Product, len(src))

	for i := range src {
		res[i] = repackOfferToProto(src[i])
	}

	return res
}

func repackParamToProto(src models.Param) *gen.Product_Param {
	valueStr := &structpb.Struct{}
	_ = protojson.Unmarshal([]byte(src.Value), valueStr)

	return &gen.Product_Param{
		Name:  src.Name,
		Type:  src.Type,
		Value: valueStr,
	}
}

func repackParamsToProto(src []models.Param) []*gen.Product_Param {
	res := make([]*gen.Product_Param, len(src))

	for i := range src {
		res[i] = repackParamToProto(src[i])
	}

	return res
}

func repackOfferToProto(src models.Offer) *gen.Product {
	price, _ := strconv.Atoi(src.Price)

	return &gen.Product{
		ID:          int32(src.ID),
		Name:        src.Name,
		Available:   src.Available,
		Type:        src.Type,
		Url:         src.URL,
		Price:       int32(price),
		Vendor:      src.Vendor,
		Description: src.Description,
		Params:      repackParamsToProto(src.Param),
	}
}

func (s *PimService) searchOffers(ctx context.Context, tx *sqlx.Tx, shopID int, searchTerm string, limit, offset int) ([]models.Offer, error) {
	const query = `
		select
			id, item_id, shop_id, name,
			available, type, url, price,
			currency_code, category_id,
			vendor, description
		from product_information.products,
			to_tsvector('russian', name) as src,
			plainto_tsquery('russian', $2) as query
		where shop_id = $1 and src @@ query
		order by ts_rank(src, query) desc
		limit $3
		offset $4;
	`

	var offers []models.Offer

	err := s.db.InTx(ctx, nil, func(tx *sqlx.Tx) error {
		rows, err := tx.Query(query, shopID, searchTerm, limit, offset)
		if err != nil {
			return err
		}
		defer rows.Close()

		for rows.Next() {
			var offer models.Offer

			err = rows.Scan(
				&offer.ID,
				&offer.ItemID,
				&offer.ShopID,
				&offer.Name,
				&offer.Available,
				&offer.Type,
				&offer.URL,
				&offer.Price,
				&offer.CurrencyID,
				&offer.CategoryID,
				&offer.Vendor,
				&offer.Description,
			)
			if err != nil {
				return err
			}

			offers = append(offers, offer)
		}

		return nil
	})

	return offers, err
}

func (s *PimService) getOffersParamsByIDs(ctx context.Context, tx *sqlx.Tx, offersIDs []int) (map[int][]models.Param, error) {
	const query = `
		select product_id as offer_id, name, type, value
		from product_information.products_attributes
		where product_id = any($1)
	`

	type param struct {
		OfferID int
		models.Param
	}

	var params []param

	err := s.db.InTx(ctx, tx, func(tx *sqlx.Tx) error {
		rows, err := tx.Query(query, pq.Array(offersIDs))
		if err != nil {
			return err
		}
		defer rows.Close()

		for rows.Next() {
			var item param

			err = rows.Scan(
				&item.OfferID,
				&item.Name,
				&item.Type,
				&item.Value,
			)
			if err != nil {
				return err
			}

			params = append(params, item)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	res := make(map[int][]models.Param)

	for _, param := range params {
		res[param.OfferID] = append(res[param.OfferID], param.Param)
	}

	return res, nil
}
