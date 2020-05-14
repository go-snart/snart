package route

import (
	"fmt"
	dg "github.com/bwmarrin/discordgo"
	re "regexp"
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

var BotOwner Okay = func(c *Ctx) bool {
	_f := "BotOwner"
	app, err := c.Session.Application("@me")
	if err != nil {
		err = fmt.Errorf("app @me: %w", err)
		Log.Error(_f, err)
		return false
	}
	return app.Owner.ID == c.Message.Author.ID
}

func UserInFunc(f func() ([]string, error)) Okay {
	_f := "UserInFunc"
	return func(c *Ctx) bool {
		ids, err := f()
		if err != nil {
			err = fmt.Errorf("f(): %w", err)
			Log.Error(_f, err)
			return false
		}

		uid := c.Message.Author.ID
		for _, id := range ids {
			if id == uid {
				return true
			}
		}

		return false
	}
}

var ServerOwner Okay = func(c *Ctx) bool {
	_f := "ServerOwner"
	gid := c.Message.GuildID
	guild, err := c.Session.Guild(gid)
	if err != nil {
		err = fmt.Errorf("guild %#v: %w", gid, err)
		Log.Error(_f, err)
		return false
	}
	return guild.OwnerID == c.Message.Author.ID
}

func SameChannelID(oc *Ctx) Okay {
	oci := oc.Message.ChannelID
	return func(c *Ctx) bool {
		return c.Message.ChannelID == oci
	}
}

func SameAuthor(oc *Ctx) Okay {
	oai := oc.Message.Author.ID
	return func(c *Ctx) bool {
		return c.Message.Author.ID == oai
	}
}

var False Okay = func(*Ctx) bool {
	return false
}

var True Okay = func(*Ctx) bool {
	return true
}

func ContentMatch(pattern string) Okay {
	_f := "ContentMatch"
	rc, err := re.Compile(pattern)
	if err != nil {
		err = fmt.Errorf("compile %#v: %w", pattern, err)
		Log.Error(_f, err)
		return False
	}

	return func(c *Ctx) bool {
		return rc.MatchString(c.Message.Content)
	}
}

var GuildAdmin Okay = func(c *Ctx) bool {
	_f := "GuildAdmin"

	perm, err := c.Session.State.UserChannelPermissions(c.Message.Author.ID, c.Message.ChannelID)
	if err != nil {
		Log.Warn(_f, err)
		return false
	}

	if perm&(dg.PermissionAdministrator|dg.PermissionManageServer) > 0 {
		return true
	}

	g, err := c.Session.State.Guild(c.Message.GuildID)
	if err != nil {
		Log.Warn(_f, err)
		return false
	}

	if c.Message.Author.ID == g.OwnerID {
		return true
	}

	return false
}
