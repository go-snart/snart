package bot

import (
	"fmt"

	"github.com/go-snart/snart/route"
)

// Admin checks if the author has bot-wide admin priveleges. (implements route.Okay)
func (b *Bot) Admin(c *route.Ctx) bool {
	_f := "(*Bot).Admin"

	app, err := c.Session.Application("@me")
	if err != nil {
		err = fmt.Errorf("app @me: %w", err)
		Log.Warn(_f, err)
	} else if app.Owner.ID == c.Message.Author.ID {
		return true
	}

	adminIDs, err := b.DB.AdminIDs()
	if err != nil {
		err = fmt.Errorf("admin ids: %w", err)
		Log.Warn(_f, err)
	} else {
		for _, adminID := range adminIDs {
			if adminID == c.Message.Author.ID {
				return true
			}
		}
	}

	return false
}
