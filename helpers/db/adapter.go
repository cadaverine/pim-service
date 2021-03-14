package db

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type adapter struct {
	db *sqlx.DB
}

// InTx ...
func (a *adapter) InTx(ctx context.Context, tx *sqlx.Tx, fn func(*sqlx.Tx) error) (err error) {
	if tx == nil {
		tx, err = a.db.BeginTxx(ctx, nil)
		if err != nil {
			return
		}

		defer func() {
			if p := recover(); p != nil {
				_ = tx.Rollback()
				panic(p)
			} else if err != nil {
				_ = tx.Rollback()
			} else {
				err = tx.Commit()
			}
		}()
	}

	err = fn(tx)
	return err
}

func newAdapter(ctx context.Context, conf Conf) (*adapter, error) {
	db, err := sqlx.Connect("postgres", getDSN(conf))
	if err != nil {
		return nil, err
	}

	return &adapter{
		db: db,
	}, nil
}

func getDSN(conf Conf) string {
	dsnTemplate := "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"
	return fmt.Sprintf(dsnTemplate, conf.Host, conf.Port, conf.User, conf.Pass, conf.Name)
}
