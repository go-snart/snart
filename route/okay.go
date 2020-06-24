package route

import (
	"fmt"

	dg "github.com/bwmarrin/discordgo"
)

type Okay func(*Ctx) bool

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

var False Okay = func(*Ctx) bool {
	return false
}

var True Okay = func(*Ctx) bool {
	return true
}

var GuildAdmin Okay = func(c *Ctx) bool {
	_f := "GuildAdmin"

	perm, err := c.Session.UserChannelPermissions(c.Message.Author.ID, c.Message.ChannelID)
	if err != nil {
		err = fmt.Errorf("perm: %w", err)
		Log.Error(_f, err)
		return false
	}

	return perm&(dg.PermissionAdministrator|
		dg.PermissionManageServer) > 0
}
