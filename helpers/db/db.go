package db

import (
	"context"

	"github.com/jmoiron/sqlx"
)

// Conf config for db connection
type Conf struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}

// IAdapter db adapter interface
type IAdapter interface {
	InTx(ctx context.Context, tx *sqlx.Tx, fn func(*sqlx.Tx) error) (err error)
}

// New creates new db adapter
func New(ctx context.Context, useMock bool, conf Conf) (IAdapter, error) {
	if useMock {
		return newMock()
	}

	return newAdapter(ctx, conf)
}
