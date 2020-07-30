// Package db contains the workings of a DB for a Snart Bot.
package db

import (
	"context"
	"errors"
	"fmt"

	pgx "github.com/jackc/pgx/v4"
	"github.com/superloach/minori"
)

// ErrNoConfigs occurs when no good configs are found.
var ErrNoConfigs = errors.New("no good configs found")

// Log is the logger for the db package.
var Log = minori.GetLogger("db")

// DB wraps a PostgreSQL connection.
type DB struct {
	Configs []*pgx.ConnConfig
}

// New creates a database abstraction.
func New() (*DB, error) {
	const _f = "New"

	sconfs := Configs()
	confs := []*pgx.ConnConfig(nil)

	for _, sconf := range sconfs {
		conf, err := pgx.ParseConfig(sconf)
		if err != nil {
			err = fmt.Errorf("parse %q: %w", sconf, err)

			Log.Warn(_f, err)
		} else {
			confs = append(confs, conf)
		}
	}

	if len(confs) == 0 {
		return nil, ErrNoConfigs
	}

	return &DB{
		Configs: confs,
	}, nil
}

// ConnKey is the context key type used by (*DB).Conn.
type ConnKey struct{}

// Conn retrieves a PostgreSQL connection for a given context, inserting the new value if necessary.
func (d *DB) Conn(ctx *context.Context) *pgx.Conn {
	const _f = "(*DB).Conn"

	val, ok := (*ctx).Value(ConnKey{}).(*pgx.Conn)
	if ok && val != nil {
		return val
	}

	for _, conf := range d.Configs {
		conn, err := pgx.ConnectConfig(*ctx, conf)
		if err != nil {
			err = fmt.Errorf("connect %q: %w", conf.ConnString(), err)
			Log.Warn(_f, err)
		} else {
			*ctx = context.WithValue(*ctx, ConnKey{}, conn)

			return conn
		}
	}

	Log.Fatal(_f, "unable to open a connection")

	return nil
}
