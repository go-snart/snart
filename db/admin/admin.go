// Package admin handles bot-wide administration permissions.
package admin

import (
	"fmt"

	"github.com/go-snart/snart/db"
	"github.com/go-snart/snart/logs"
	"github.com/go-snart/snart/route"
)

// IsAdmin checks if the author has bot-wide admin privileges.
func IsAdmin(d *db.DB) route.Okay {
	return func(c *route.Ctx) bool {
		admins, err := List(d)
		if err != nil {
			err = fmt.Errorf("list admins: %w", err)
			logs.Warn.Println(err)

			return false
		}

		for _, admin := range admins {
			if c.Message.Author.ID == admin {
				return true
			}
		}

		app, err := c.Session.Application("@me")
		if err != nil {
			err = fmt.Errorf("app @me: %w", err)
			logs.Warn.Println(err)

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
func List(d *db.DB) ([]string, error) {
	return nil, fmt.Errorf("stub")
}
