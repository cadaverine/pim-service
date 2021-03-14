package db

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type mock struct{}

func (m *mock) InTx(ctx context.Context, tx *sqlx.Tx, fn func(*sqlx.Tx) error) (err error) {
	return nil
}

func newMock() (*mock, error) {
	return &mock{}, nil
}
