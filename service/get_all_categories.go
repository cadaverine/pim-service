package service

import (
	"context"

	gen "gitlab.com/cadaverine/pim-service/gen/go/api/pim-service"
)

func (s *PimService) GetAllCategories(ctx context.Context, req *gen.Empty) (*gen.Categories, error) {
	return nil, nil
}
