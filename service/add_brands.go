package service

import (
	"context"

	"github.com/jackc/pgx/v4"
	gen "gitlab.com/cadaverine/pim-service/gen"
)

// AddBrands ...
func (s *PimService) AddBrands(ctx context.Context, req *gen.Brands) (*gen.Empty, error) {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	for _, brand := range req.GetBrands() {
		err = addBrand(ctx, tx, brand)
		if err != nil {
			return nil, err
		}
	}

	return &gen.Empty{}, nil
}

func addBrand(ctx context.Context, tx pgx.Tx, brand *gen.Brand) error {
	const query = `
		insert into catalogs.brands (name, logo, site, company)
		values ($1, $2, $3, $4);
	`

	_, err := tx.Exec(ctx, query,
		brand.GetName(),
		brand.GetLogo(),
		brand.GetSite(),
		brand.GetCompany(),
	)
	if err != nil {
		return err
	}

	return nil
}
