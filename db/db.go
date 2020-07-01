package db

import (
	"fmt"

	"github.com/superloach/minori"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

type DB struct {
	*r.Session

	Host string
	Port int
	User string
	Pass string

	Cache *MapCache

	failed bool
}

var Log = minori.GetLogger("db")

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
