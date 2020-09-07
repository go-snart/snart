// Package admin is a Snart plug which provides basic administration commands.
package admin

import (
	"github.com/go-snart/snart/bot/plug"
	"github.com/go-snart/snart/route"
)

var Plug = plug.Plug(&Admin{})

type Admin struct {
	plug.Base
}

func (a *Admin) String() string {
	return "admin"
}

func (a *Admin) PlugHandler(r *route.Handler) {
	a.Base.PlugHandler(r)

	r.Add(
		&route.Route{
			Name:  "restart",
			Match: route.MustMatch("restart"),
			Desc:  "restart the bot",
			Okay:  route.BotAdmin(a.DB),
			Cat:   a.String(),
			Func:  a.Restart,
		},
		&route.Route{
			Name:  "prefix",
			Match: route.MustMatch("prefix|pfx"),
			Desc:  "changes prefixes",
			Okay:  nil,
			Cat:   a.String(),
			Func:  a.Prefix,
		},
	)
}
