package db

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type mock struct{}

func (m *mock) InTx(ctx context.Context, tx pgx.Tx, fn func(pgx.Tx) error) (err error) {
	return nil
}

func newMock() (*mock, error) {
	return &mock{}, nil
}
