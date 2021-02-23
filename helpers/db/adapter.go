package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/log15adapter"
	"github.com/jackc/pgx/v4/pgxpool"

	log "gopkg.in/inconshreveable/log15.v2"
)

type adapter struct {
	db *pgxpool.Pool
}

// InTx ...
func (a *adapter) InTx(ctx context.Context, tx pgx.Tx, fn func(pgx.Tx) error) (err error) {
	if tx == nil {
		tx, err = a.db.Begin(ctx)
		if err != nil {
			return
		}

		defer func() {
			if p := recover(); p != nil {
				_ = tx.Rollback(ctx)
				panic(p)
			} else if err != nil {
				_ = tx.Rollback(ctx)
			} else {
				err = tx.Commit(ctx)
			}
		}()
	}

	err = fn(tx)
	return err
}

func newAdapter(ctx context.Context, conf Conf) (*adapter, error) {
	poolConf, err := getPoolConf(conf)
	if err != nil {
		return nil, err
	}

	db, err := pgxpool.ConnectConfig(ctx, poolConf)
	if err != nil {
		return nil, err
	}

	return &adapter{
		db: db,
	}, nil
}

func getPoolConf(conf Conf) (*pgxpool.Config, error) {
	dsnTemplate := "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"

	dsn := fmt.Sprintf(dsnTemplate, conf.Host, conf.Port, conf.User, conf.Pass, conf.Name)

	logger := log15adapter.NewLogger(log.New("module", "pgx"))

	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	poolConfig.ConnConfig.Logger = logger

	return poolConfig, nil
}
