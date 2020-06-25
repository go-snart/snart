package db

import (
	"fmt"

	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

type Admin struct {
	ID string `rethinkdb:"id"`
}

var AdminTable = r.DB("config").TableCreate("admin")

func (d *DB) AdminIDs() ([]string, error) {
	_f := "(*DB).AdminIDs"

	d.Once(ConfigDB)
	d.Once(AdminTable)

	admins := make([]Admin, 0)
	q := r.DB("config").Table("admin")

	err := q.ReadAll(&admins, d)
	if err != nil {
		err = fmt.Errorf("readall &admins: %w", err)
		Log.Error(_f, err)
		return nil, err
	}

	adminIDs := make([]string, len(admins))
	for i, admin := range admins {
		adminIDs[i] = admin.ID
	}

	return adminIDs, nil
}
