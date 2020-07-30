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
	const (
		_f = "Table"
		e  = `CREATE TABLE IF NOT EXISTS admin(
			id TEXT PRIMARY KEY UNIQUE
		)`
	)

	_, err := d.Conn(&ctx).Exec(ctx, e)
	if err != nil {
		err = fmt.Errorf("exec %#q: %w", e, err)

		Log.Error(_f, err)

		return
	}
}

// IsAdmin checks if the author has bot-wide admin privileges.
func IsAdmin(d *db.DB) route.Okay {
	return func(c *route.Ctx) bool {
		const _f = "IsAdmin"

		for _, admin := range List(c, d) {
			if c.Message.Author.ID == admin {
				return true
			}
		}

		app, err := c.Session.Application("@me")
		if err != nil {
			err = fmt.Errorf("app @me: %w", err)
			Log.Warn(_f, err)

			return false
		}

		if app.Owner != nil && c.Message.Author.ID == app.Owner.ID {
			return true
		}

		if app.Team != nil {
			if c.Message.Author.ID == app.Team.OwnerID {
				return true
			}

			for _, member := range app.Team.Members {
				if c.Message.Author.ID == member.User.ID {
					return true
				}
			}
		}

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
		err = fmt.Errorf("query %#q: %w", q, err)

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
