// Package admin handles bot-wide administration permissions.
package admin

import (
	"context"
	"fmt"

	"github.com/go-snart/snart/db"
	"github.com/go-snart/snart/route"
)

// Table builds the table of admins.
func Table(ctx context.Context, d *db.DB) {
	const q = `CREATE TABLE IF NOT EXISTS admin(id TEXT PRIMARY KEY UNIQUE)`
	x, err := d.Conn(&ctx).Exec(ctx, q)
	Log.Debugf("table", "%#v %#v", x, err)
}

// IsAdmin checks if the author has bot-wide admin privileges.
func IsAdmin(d *db.DB) route.Okay {
	return func(c *route.Ctx) bool {
		return false
	}
}

// List returns a list of known admin IDs from the database.
func List(ctx context.Context, d *db.DB) []string {
	const _f = "List"

	Table(ctx, d)

	const q = `SELECT id FROM admin`

	rows, err := d.Conn(&ctx).Query(ctx, q)
	if err != nil {
		err = fmt.Errorf("db query %#q: %w", q, err)
		Log.Warn(_f, err)

		return nil
	}
	defer rows.Close()

	list := []string(nil)

	for rows.Next() {
		admin := ""

		err = rows.Scan(&admin)
		if err != nil {
			err = fmt.Errorf("scan admin: %w", err)
			Log.Warn(_f, err)

			return nil
		}

		list = append(list, admin)
	}

	err = rows.Err()
	if err != nil {
		err = fmt.Errorf("rows err: %w", err)
		Log.Warn(_f, err)

		return nil
	}

	return list
}
