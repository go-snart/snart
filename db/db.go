// Package db contains the workings of a DB for a Snart Bot.
package db

import (
	"fmt"

	"github.com/superloach/minori"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

// Log is the logger for the db package.
var Log = minori.GetLogger("db")

// DB wraps a RethinkDB session and Cache.
type DB struct {
	*r.Session

	Host string
	Port int
	User string
	Pass string

	Cache *MapCache

	failed bool
}

// Start performs a DB's startup functions.
func (d *DB) Start() error {
	_f := "(*DB).Start"

	s, err := r.Connect(r.ConnectOpts{
		Address:  fmt.Sprintf("%s:%d", d.Host, d.Port),
		Username: d.User,
		Password: d.Pass,
	})
	if err != nil {
		d.failed = true

		err = fmt.Errorf("connect %s:%s@%s:%d: %w", d.User, d.Pass, d.Host, d.Port, err)
		Log.Error(_f, err)
		return err
	}
	d.Session = s

	d.Cache = NewMapCache()

	return nil
}
