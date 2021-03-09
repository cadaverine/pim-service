package service

import (
	"context"

	"github.com/jackc/pgx/v4"
	gen "gitlab.com/cadaverine/pim-service/gen/go/api/pim-service"
)

// AddBrands ...
func (s *PimService) AddBrands(ctx context.Context, req *gen.Brands) (*gen.Empty, error) {
	const query = `
		insert into catalogs.brands (name, logo, site, company)
		values ($1, $2, $3, $4);
	`

	batch := &pgx.Batch{}

	for _, brand := range req.GetBrands() {
		batch.Queue(query,
			brand.GetName(),
			brand.GetLogo(),
			brand.GetSite(),
			brand.GetCompany(),
		)
	}

	return &gen.Empty{}, s.db.InTx(ctx, nil, func(tx pgx.Tx) error {
		_, err := tx.SendBatch(ctx, batch).Exec()
		return err
	})
}
