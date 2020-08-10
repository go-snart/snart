// Package db contains the workings of a DB for a Snart Bot.
package db

import (
	"context"
	"fmt"

	pgx "github.com/jackc/pgx/v4"

	"github.com/go-snart/snart/logs"
)

const _p = "snart/db"

// DB wraps a PostgreSQL connection.
type DB struct {
	Configs []*pgx.ConnConfig
}

// New creates a database abstraction.
func New() *DB {
	sconfs := Configs()
	confs := []*pgx.ConnConfig(nil)

	for _, sconf := range sconfs {
		conf, err := pgx.ParseConfig(sconf)
		if err != nil {
			err = fmt.Errorf("parse %q: %w", sconf, err)

			logs.Warn.Println(err)
		} else {
			confs = append(confs, conf)
		}
	}

	if len(confs) == 0 {
		logs.Info.Fatalln("no good configs found")

		return nil
	}

	return &DB{
		Configs: confs,
	}
}

// ConnKey is the context key type used by (*DB).Conn.
type ConnKey struct{}

// Conn retrieves a PostgreSQL connection for a given context, inserting the new value if necessary.
func (d *DB) Conn(ctx *context.Context) *pgx.Conn {
	val, ok := (*ctx).Value(ConnKey{}).(*pgx.Conn)
	if ok && val != nil {
		return val
	}

	for _, conf := range d.Configs {
		conn, err := pgx.ConnectConfig(*ctx, conf)
		if err != nil {
			err = fmt.Errorf("connect %q: %w", conf.ConnString(), err)
			logs.Warn.Println(err)
		} else {
			*ctx = context.WithValue(*ctx, ConnKey{}, conn)

			return conn
		}
	}

	logs.Info.Fatalln("unable to open a connection")

	return nil
}
