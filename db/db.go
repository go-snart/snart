// Package db contains the workings of a DB for a Snart Bot.
package db

import (
	"context"
	"fmt"

	pgx "github.com/jackc/pgx/v4"
	"github.com/superloach/minori"

	"github.com/go-snart/snart/db/cache"
)

// Log is the logger for the db package.
var Log = minori.GetLogger("db")

// DB wraps a RethinkDB session and Cache.
type DB struct {
	*pgx.Conn
	*pgx.ConnConfig

	Cache cache.Cache
}

func NewDB(config string) (*DB, error) {
	conf, err := pgx.ParseConfig(config)
	if err != nil {
		return nil, err
	}

	return &DB{
		Conn: nil,
		ConnConfig: conf,
		Cache: cache.NewMapCache(),
	}, nil
}

// Start performs a DB's startup functions.
func (d *DB) Start(ctx context.Context) error {
	_f := "(*DB).Start"

	Log.Debugf(_f, "%#v", d)

	conn, err := pgx.ConnectConfig(ctx, d.ConnConfig)
	if err != nil {
		err = fmt.Errorf("connect %#v: %w", d.ConnConfig, err)
		Log.Error(_f, err)

		return err
	}

	d.Conn = conn

	return nil
}
