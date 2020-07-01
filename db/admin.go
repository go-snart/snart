package db

import (
	"fmt"

	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

// Admin contains the user ID of an administrator.
type Admin struct {
	ID string `rethinkdb:"id"`
}

// AdminTable is a table builder for config.admin.
var AdminTable = BuildTable(
	ConfigDB, "admin",
	nil, nil,
)

// AdminIDs is a convenience method for fetching a list of administrators.
func (d *DB) AdminIDs() ([]string, error) {
	_f := "(*DB).AdminIDs"

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
