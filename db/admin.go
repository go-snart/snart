package db

import (
	"fmt"

	"github.com/go-snart/snart/route"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

// Admin contains the user ID of an administrator.
type Admin struct {
	ID string `rethinkdb:"id"`
}

// AdminTable is a table builder for config.admin.
var AdminTable = BuildTable(ConfigDB, "admin")

// Admin checks if the author has bot-wide admin privileges (implements route.Okay).
func (d *DB) Admin(c *route.Ctx) bool {
	_f := "(*DB).Admin"

	Log.Debug(_f, "admin")

	Log.Debug(_f, "lock d.cache")
	d.Cache.Lock()

	Log.Debug(_f, "get admin")
	admin := d.Cache.Get("admin").(Cache)

	Log.Debug(_f, "unlock d.cache")
	d.Cache.Unlock()

	Log.Debug(_f, "lock admin")
	admin.Lock()
	defer admin.Unlock()

	Log.Debug(_f, "has?")
	return admin.Has(c.Message.Author.ID)
}

// NeighCache maintains a running state of known Neighs.
func (d *DB) AdminCache() {
	_f := "(*DB).AdminCache"

	d.WaitReady()

	q := AdminTable.Build(d).Changes(
		r.ChangesOpts{IncludeInitial: true},
	)

	curs, err := q.Run(d)
	if err != nil {
		err = fmt.Errorf("db run %s: %w", q, err)
		Log.Error(_f, err)

		return
	}
	defer curs.Close()

	chng := struct {
		New *Admin `rethinkdb:"new_val"`
		Old *Admin `rethinkdb:"old_val"`
	}{}

	Log.Debug(_f, "lock d.cache")
	d.Cache.Lock()

	Log.Debug(_f, "!has?")
	if !d.Cache.Has("admin") {
		Log.Debug(_f, "new cache")
		d.Cache.Set("admin", NewMapCache())
	}

	Log.Debug(_f, "get admin")
	admin := d.Cache.Get("admin").(Cache)

	Log.Debug(_f, "unlock d.cache")
	d.Cache.Unlock()

	for curs.Next(&chng) {
		Log.Debug(_f, "lock admin")
		admin.Lock()

		if chng.New != nil {
			admin.Set(chng.New.ID, chng.New)
		} else {
			admin.Del(chng.Old.ID)
		}

		Log.Debug(_f, "unlock admin")
		admin.Unlock()
	}

	if err := curs.Err(); err != nil {
		resp, ok := curs.NextResponse()

		err = fmt.Errorf(
			"curs err: %w\n"+
				"chng is %#v/%#v\n"+
				"resp(%v) is %q",
			err,
			chng.New, chng.Old,
			ok, resp,
		)
		Log.Error(_f, err)

		return
	}
}
