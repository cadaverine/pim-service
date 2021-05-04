package service

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"gitlab.com/cadaverine/pim-service/models"
)

// SaveProducts ...
func (s *PimService) SaveProducts(ctx context.Context, tx *sqlx.Tx, shopID int, offers []models.Offer) error {
	return s.db.InTx(ctx, tx, func(t *sqlx.Tx) error {
		err := s.deleteProducts(ctx, tx, shopID)
		if err != nil {
			return errors.Wrap(err, "failed to delete products")
		}

		err = s.addProducts(ctx, tx, shopID, offers)
		if err != nil {
			return errors.Wrapf(err, "failed to add products for shop '%v'", shopID)
		}

		return nil
	})
}

func (s *PimService) addProduct(ctx context.Context, tx *sqlx.Tx, shopID int, offer models.Offer) (int, error) {
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
		) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) on conflict (shop_id, item_id) do
		update set
			item_id = excluded.item_id,
			shop_id = excluded.shop_id,
			name = excluded.name,
			available = excluded.available,
			type = excluded.type,
			url = excluded.url,
			price = excluded.price,
			currency_code = excluded.currency_code,
			vendor = excluded.vendor,
			description = excluded.description,
			deleted_at = null
		returning id;
	`

	var productID int

	price, _ := strconv.Atoi(offer.Price)

	err := s.db.InTx(ctx, tx, func(tx *sqlx.Tx) error {
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
			return errors.Wrapf(err, "failed to save product: $+v", offer)
		}
		return nil
	})

	return productID, err
}

func (s *PimService) addProductAttributes(ctx context.Context, tx *sqlx.Tx, productID int, params []models.Param) error {
	const (
		defaultType = "string"
		query       = `
			insert into product_information.products_attributes (
				product_id,
				name,
				type,
				value
			) values ($1,$2,$3,$4) on conflict (product_id, name) do
			update set
				product_id = excluded.product_id,
				name = excluded.name,
				type = excluded.type,
				value = excluded.value;
		`
	)

	return s.db.InTx(ctx, tx, func(tx *sqlx.Tx) error {
		for i := range params {
			value, _ := json.Marshal(map[string]interface{}{
				"type":  defaultType,
				"value": params[i].Value,
			})

			_, err := tx.Exec(query,
				productID,
				params[i].Name,
				defaultType,
				value,
			)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (s *PimService) addProductCategory(ctx context.Context, tx *sqlx.Tx, productID, categoryID int) error {
	const query = `
		insert into product_information.products_categories (product_id, category_id)
		values ($1, $2);
	`

	return s.db.InTx(ctx, tx, func(tx *sqlx.Tx) error {
		_, err := tx.Exec(query, productID, categoryID)
		return err
	})
}

func (s *PimService) getCategoryID(ctx context.Context, tx *sqlx.Tx, shopID int, categoryItemID string) (int, error) {
	const query = `
		select id
		from product_information.categories c
		where shop_id = $1 and item_id = $2
	`

	var id int

	return id, s.db.InTx(ctx, tx, func(tx *sqlx.Tx) error {
		return tx.Get(&id, query, shopID, categoryItemID)
	})
}

func (s *PimService) addProducts(ctx context.Context, tx *sqlx.Tx, shopID int, offers []models.Offer) error {
	return s.db.InTx(ctx, tx, func(tx *sqlx.Tx) error {
		for _, offer := range offers {
			productID, err := s.addProduct(ctx, tx, shopID, offer)
			if err != nil {
				return err
			}

			categoryID, err := s.getCategoryID(ctx, tx, shopID, offer.CategoryID)
			if err != nil {
				return err
			}

			err = s.addProductCategory(ctx, tx, productID, categoryID)
			if err != nil {
				return err
			}

			err = s.addProductAttributes(ctx, tx, productID, offer.Param)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (s *PimService) deleteProducts(ctx context.Context, tx *sqlx.Tx, shopID int) error {
	const query = `
		update product_information.products
		set deleted_at = now()
		where shop_id = $1;
	`

	return s.db.InTx(ctx, tx, func(tx *sqlx.Tx) error {
		_, err := tx.Exec(query, shopID)
		return err
	})
}
