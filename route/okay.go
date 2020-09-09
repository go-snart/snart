package route

import (
	"fmt"

	dg "github.com/bwmarrin/discordgo"

	"github.com/go-snart/snart/db"
	"github.com/go-snart/snart/log"
)

// Okay is a function which checks if a Ctx should be used.
type Okay func(*Ctx) bool

// Any is a logical OR of Okays.
func Any(chs ...Okay) Okay {
	return func(c *Ctx) bool {
		for _, ch := range chs {
			if ch(c) {
				return true
			}
		}

		return false
	}
}

// All is a logical AND of Okays.
func All(chs ...Okay) Okay {
	return func(c *Ctx) bool {
		for _, ch := range chs {
			if !ch(c) {
				return false
			}
		}

		return true
	}
}

// False is an Okay that always returns false.
var False Okay = func(*Ctx) bool {
	return false
}

// True is an Okay that always returns true.
var True Okay = func(*Ctx) bool {
	return true
}

// GuildAdmin is an Okay that checks if the user has administrator privileges on the guild.
var GuildAdmin Okay = func(c *Ctx) bool {
	perm, err := c.Session.UserChannelPermissions(c.Message.Author.ID, c.Message.ChannelID)
	if err != nil {
		err = fmt.Errorf("perm: %w", err)

		log.Warn.Println(err)

		return false
	}

	return perm&(dg.PermissionAdministrator|
		dg.PermissionManageServer) > 0
}

// BotAdmin is a route.Okay that checks if the author has bot-wide admin privileges.
func BotAdmin(d *db.DB) Okay {
	return func(c *Ctx) bool {
		admins, err := d.Admins(c)
		if err != nil {
			err = fmt.Errorf("list admins: %w", err)
			log.Warn.Println(err)

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
			log.Warn.Println(err)

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
