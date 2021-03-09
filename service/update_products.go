package service

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

// UpdateProducts ...
func (s *PimService) UpdateProducts(ctx context.Context, tx pgx.Tx, shop string, offers []offer) error {
	return s.db.InTx(ctx, tx, func(t pgx.Tx) error {
		err := s.deleteProducts(ctx, tx)
		if err != nil {
			return errors.Wrap(err, "failed to delete products")
		}

		err = s.addProducts(ctx, tx, shop, offers)
		if err != nil {
			return errors.Wrapf(err, "failed to add categories for shop '%s'", shop)
		}

		return nil
	})
}

func (s *PimService) addProducts(ctx context.Context, tx pgx.Tx, shop string, offers []offer) error {
	const query = `
		insert into products.items (
			id, name, available,
			type, url, price, currency_code,
			vendor, description
		) values ($1,$2,$3,$4,$5,$6,$7,$8,$9) on conflict (id) do
		update set
			id = excluded.id,
			name = excluded.name,
			available = excluded.available,
			type = excluded.type,
			url = excluded.url,
			price = excluded.price,
			currency_code = excluded.currency_code,
			vendor = excluded.vendor,
			description = excluded.description,
			deleted_at = null;
	`

	batch := &pgx.Batch{}

	for _, offer := range offers {
		batch.Queue(query,
			offer.ID,
			offer.Name,
			offer.Available,
			offer.Type,
			offer.URL,
			offer.Price,
			offer.CurrencyID,
			offer.Vendor,
			offer.Description,
		)
	}

	return s.db.InTx(ctx, tx, func(t pgx.Tx) error {
		br := tx.SendBatch(ctx, batch)
		defer br.Close()

		for i := 0; i < batch.Len(); i++ {
			_, err := br.Exec()
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (s *PimService) deleteProducts(ctx context.Context, tx pgx.Tx) error {
	const query = `
		update products.items
		set deleted_at = now();
	`

	return s.db.InTx(ctx, tx, func(t pgx.Tx) error {
		_, err := tx.Exec(ctx, query)
		return err
	})
}
