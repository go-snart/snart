package route_test

import (
	"fmt"
	"os"

	dg "github.com/bwmarrin/discordgo"
	"github.com/go-snart/snart/route"
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

	return msg, &dg.MessageCreate{
		Message: msg,
	}
}

func sessionDummy() (
	string,
	*dg.Session,
) {
	tok := os.Getenv("SNART_TOKEN")
	if tok == "" {
		panic("please provide SNART_TOKEN")
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
	string, string, *dg.Session, *dg.Message, *route.Flags, *route.Route,
	*route.Ctx,
) {
	var (
		prefix      = "./"
		cleanPrefix = "./"
		flags       = &route.Flags{}
	)

	_, _, _, _,
		message := messageDummy(content)

	_,
		session := sessionDummy()

	_, _, _, _, _, _,
		r := routeDummy()

	return prefix, cleanPrefix, session, message, flags, r,
		&route.Ctx{
			Prefix:      prefix,
			CleanPrefix: cleanPrefix,
			Session:     session,
			Message:     message,
			Flags:       flags,
			Route:       r,
		}
}

func ctxBadDummy() *route.Ctx {
	_, _, _, _, _, _,
		c := ctxDummy("")
	c.Session = sessionBadDummy()

	return c
}

func flagsDummy(ctx *route.Ctx) (
	string, []string,
	*route.Flags,
) {
	name := ctx.Route.Name
	args := []string{
		"foo",
		"bar",
		"baz",
	}

	return name, args,
		route.NewFlags(ctx, name, args)
}

func routeDummy() (
	string, string, string, string, route.Okay, func(*route.Ctx) error,
	*route.Route,
) {
	var (
		name  = "route"
		match = "route|yeet"
		cat   = "test"
		desc  = "a test route"
		okay  = route.True
		_func = func(c *route.Ctx) error {
			c.Route.Desc = "run"
			return nil
		}
	)

	return name, match, cat, desc, okay, _func,
		&route.Route{
			Name:  name,
			Match: match,
			Cat:   cat,
			Desc:  desc,
			Okay:  okay,
			Func:  _func,
		}
}

func routerDummy() (
	*route.Route,
	*route.Router,
) {
	router := route.NewRouter()

	_, _, _, _, _, _,
		r := routeDummy()
	router.Add(r)

	return r,
		router
}
