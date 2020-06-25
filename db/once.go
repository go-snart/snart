package db

import (
	"fmt"

	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

func (d *DB) Once(q r.Term) {
	_f := "(*DB).Once"
	d.WaitReady()

	qs := q.String()
	if d.Cache["once"].Contains(qs) {
		Log.Debugf(_f, "cache %s", qs)
		return
	}
	d.Cache["once"].Add(qs, struct{}{})

	err := q.Exec(d)
	if err != nil {
		err = fmt.Errorf("exec %s: %w", q, err)
		Log.Warn(_f, err)
		return
	}
}
