package service

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	gen "gitlab.com/cadaverine/pim-service/gen/pim-service"
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
	withParams := req.GetWithParams()

	available := req.GetFilters().GetAvailable()
	categoriesIDs := req.GetFilters().GetCategoriesIDs()

	var isAvailable sql.NullBool
	if available != nil {
		isAvailable = sql.NullBool{
			Bool:  available.GetValue(),
			Valid: true,
		}
	}

	limit, offset := int(req.GetMeta().GetLimit()), int(req.GetMeta().GetOffset())

	if limit == 0 || limit > maxLimit {
		limit = defaultLimit
	}

	var err error
	var offers []models.Offer

	err = s.db.InTx(ctx, nil, func(tx *sqlx.Tx) error {
		offers, err = s.searchOffers(ctx, tx, shopID, searchTerm, categoriesIDs, isAvailable, limit, offset)
		if err != nil {
			return err
		}

		offersIDs := make([]int, len(offers))
		for i := range offers {
			offersIDs[i] = offers[i].ID
		}

		if withParams {
			offersMap, err := s.getOffersParamsByIDs(ctx, tx, offersIDs)
			if err != nil {
				return err
			}

			for i := range offers {
				offers[i].Param = offersMap[offers[i].ID]
			}
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
		ItemID:      src.ItemID,
		Name:        src.Name,
		Available:   src.Available,
		Type:        src.Type,
		Url:         src.URL,
		Price:       int32(price),
		Vendor:      src.Vendor,
		Description: src.Description,
		CurrencyID:  src.CurrencyID,
		Params:      repackParamsToProto(src.Param),
	}
}

func (s *PimService) searchOffers(ctx context.Context, tx *sqlx.Tx, shopID int, searchTerm string, categoriesIDs []int32, isAvailable sql.NullBool, limit, offset int) ([]models.Offer, error) {
	const query = `
		with recursive categories_cte as (
			select item_id, id
			from product_information.categories c
			where
				shop_id = $1 and
				(array_length($2::int[], 1) is null or item_id = any($2::int[]))

			union

			select c.item_id, c.id
			from product_information.categories c
			join categories_cte cte on cte.id = c.parent_id
			where shop_id = $1
		)
		select
			product.id, product.item_id, product.shop_id, product.name,
			product.available, product.type, product.url, product.price,
			product.currency_code, product.vendor, product.description
		from product_information.products product
		join product_information.products_categories pc on pc.product_id = product.id
		join categories_cte category on category.id = pc.category_id
		where shop_id = $1 and ($4::bool is null or available = $4)
		order by name <-> $3
		limit $5
		offset $6;
	`

	var offers []models.Offer

	err := s.db.InTx(ctx, nil, func(tx *sqlx.Tx) error {
		rows, err := tx.Query(query, shopID, pq.Int32Array(categoriesIDs), searchTerm, isAvailable, limit, offset)
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
