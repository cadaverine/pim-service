package service

import (
	"context"

	"github.com/jmoiron/sqlx"
	"gitlab.com/cadaverine/pim-service/models"
)

// SaveShop ...
func (s *PimService) SaveShop(ctx context.Context, shop models.Shop) (int, error) {
	const query = `
		insert into product_information.shops (name, company, url, platform)
		values ($1, $2, $3, $4) on conflict (name, company) do
		update set
			name = excluded.name,
			company = excluded.company,
			url = excluded.url,
			platform = excluded.platform
		returning id;
	`

	var shopID int

	err := s.db.InTx(ctx, nil, func(tx *sqlx.Tx) error {
		return tx.Get(&shopID, query,
			shop.Name,
			shop.Company,
			shop.URL,
			shop.Platform,
		)
	})

	return shopID, err
}
