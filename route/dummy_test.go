package route_test

import (
	"context"
	"fmt"

	dg "github.com/bwmarrin/discordgo"

	"github.com/go-snart/snart/db"
	"github.com/go-snart/snart/db/prefix"
	"github.com/go-snart/snart/db/token"
	"github.com/go-snart/snart/logs"
	"github.com/go-snart/snart/route"
)

func val(v interface{}) string {
	return fmt.Sprintf("%#v", v)
}

func prefixDummy() (
	string, string,
	*prefix.Prefix,
) {
	logs.Debug.Println(".")

	const (
		value = "./"
		clean = "./"
	)

	return value, clean,
		&prefix.Prefix{
			Value: value,
			Clean: clean,
		}
}

func messageDummy(content string) (
	string, string, string, *dg.User,
	*dg.Message,
) {
	logs.Debug.Println(".")

	const (
		id        = "12345678900"
		channelID = "12345678901"
		guildID   = "12345678902"
	)

	var author = &dg.User{}

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
	logs.Debug.Println(".")

	_, _, _, _,
		msg := messageDummy(content)

	return msg, &dg.MessageCreate{
		Message: msg,
	}
}

var sessionDummy = token.Open(context.Background(), db.New(), nil)

var sessionBadDummy = func() *dg.Session {
	logs.Debug.Println(".")

	session, _ := dg.New("foo")

	return session
}()

func ctxDummy(content string) (
	*prefix.Prefix, *dg.Session, *dg.Message, *route.Flag, *route.Route,
	*route.Ctx,
) {
	logs.Debug.Println(".")

	flag := &route.Flag{}
	session := sessionDummy

	_, _,
		pfx := prefixDummy()

	_, _, _, _,
		message := messageDummy(content)

	_, _, _, _, _, _,
		r := routeDummy()

	return pfx, session, message, flag, r,
		&route.Ctx{
			Prefix:  pfx,
			Session: session,
			Message: message,
			Flag:    flag,
			Route:   r,
		}
}

func ctxBadDummy() *route.Ctx {
	_, _, _, _, _,
		c := ctxDummy("")

	ses := sessionBadDummy

	c.Session = ses

	return c
}

func flagDummy(ctx *route.Ctx) (
	string, []string,
	*route.Flag,
) {
	name := ctx.Route.Name
	args := []string{
		"foo",
		"bar",
		"baz",
	}

	return name, args,
		route.NewFlag(ctx, name, args)
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
