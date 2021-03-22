package service

import (
	"gitlab.com/cadaverine/pim-service/helpers/db"

	gen "gitlab.com/cadaverine/pim-service/gen/pim-service"
)

// PimService ...
type PimService struct {
	db db.IAdapter
	gen.UnimplementedPimServiceServer
}

// NewPimService ...
func NewPimService(db db.IAdapter) *PimService {
	return &PimService{db: db}
}
