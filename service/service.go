package service

import "github.com/jackc/pgx/v4/pgxpool"

// PimService ...
type PimService struct {
	db *pgxpool.Pool
}

// NewPimService ...
func NewPimService(db *pgxpool.Pool) *PimService {
	return &PimService{db}
}
