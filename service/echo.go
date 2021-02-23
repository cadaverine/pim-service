package service

import (
	"context"

	gen "gitlab.com/cadaverine/pim-service/gen/go/api/pim-service"
)

// Echo ...
func (s *PimService) Echo(ctx context.Context, req *gen.String) (*gen.String, error) {
	return req, nil
}
