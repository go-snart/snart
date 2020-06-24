package route

import (
	"fmt"
	"os"

	dg "github.com/bwmarrin/discordgo"
)

func val(v interface{}) string {
	return fmt.Sprintf("%#v", v)
}

func messageDummy(content string) (
	string, string, string, *dg.User,
	*dg.Message,
) {
	var (
		id        = "12345678900"
		channelID = "12345678901"
		guildID   = "12345678902"
		author    = &dg.User{}
	)

	return id, channelID, guildID, author,
		&dg.Message{
			ID:        id,
			ChannelID: channelID,
			GuildID:   guildID,
			Content:   content,
			Author:    author,
		}
}

func messageCreateDummy(content string) (
	*dg.Message,
	*dg.MessageCreate,
) {
	_, _, _, _,
		msg := messageDummy(content)
	return msg, &dg.MessageCreate{msg}
}

func sessionDummy() (
	string,
	*dg.Session,
) {
	tok := os.Getenv("SNART_TEST_TOKEN")
	if tok == "" {
		panic("please provide $SNART_TEST_TOKEN")
	}

	session, err := dg.New(tok)
	if err != nil {
		panic(err)
	}

	return tok, session
}

func sessionBadDummy() *dg.Session {
	session, err := dg.New("foo")
	if err != nil {
		panic(err)
	}

	return session
}

func ctxDummy(content string) (
	string, string, *dg.Session, *dg.Message, *Flags, *Route,
	*Ctx,
) {
	var (
		prefix      = "./"
		cleanPrefix = "./"
		flags       = &Flags{}
	)

	_, _, _, _,
		message := messageDummy(content)

	_,
		session := sessionDummy()

	_, _, _, _, _, _,
		route := routeDummy()

	return prefix, cleanPrefix, session, message, flags, route,
		&Ctx{
			Prefix:      prefix,
			CleanPrefix: cleanPrefix,
			Session:     session,
			Message:     message,
			Flags:       flags,
			Route:       route,
		}
}

func ctxBadDummy() *Ctx {
	_, _, _, _, _, _,
		c := ctxDummy("")
	c.Session = sessionBadDummy()
	return c
}

func flagsDummy(ctx *Ctx) (
	string, []string,
	*Flags,
) {
	name := ctx.Route.Name
	args := []string{
		"foo",
		"bar",
		"baz",
	}

	return name, args,
		NewFlags(ctx, name, args)
}

func routeDummy() (
	string, string, string, string, Okay, func(*Ctx) error,
	*Route,
) {
	var (
		name  = "route"
		match = "route|yeet"
		cat   = "test"
		desc  = "a test route"
		okay  = True
		_func = func(c *Ctx) error {
			c.Route.Desc = "run"
			return nil
		}
	)

	return name, match, cat, desc, okay, _func,
		&Route{
			Name:  name,
			Match: match,
			Cat:   cat,
			Desc:  desc,
			Okay:  okay,
			Func:  _func,
		}
}
