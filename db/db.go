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

// DB wraps a PostgreSQL connection and Cache.
type DB struct {
	*pgx.ConnConfig

	Cache cache.Cache
}

// NewDB creates a database abstraction.
func NewDB(config string) (*DB, error) {
	conf, err := pgx.ParseConfig(config)
	if err != nil {
		return nil, err
	}

	return &DB{
		ConnConfig: conf,
		Cache:      cache.NewMapCache(),
	}, nil
}

// ConnKey is the context key type used by (*DB).Conn.
type ConnKey struct{}

// Conn retrieves a PostgreSQL connection for a given context, inserting the new value if necessary.
func (d *DB) Conn(ctx *context.Context) *pgx.Conn {
	_f := "(*DB).Conn"

	val, ok := (*ctx).Value(ConnKey{}).(*pgx.Conn)
	if ok && val != nil {
		return val
	}

	conn, err := pgx.ConnectConfig(*ctx, d.ConnConfig)
	if err != nil {
		err = fmt.Errorf("connect %#v: %w", d.ConnConfig, err)
		Log.Fatal(_f, err)

		return nil
	}

	*ctx = context.WithValue(*ctx, ConnKey{}, conn)

	return conn
}
