// Package admin provides a Snart plug with basic administration features.
package admin

import (
	"github.com/go-snart/snart/bot/plug"
	"github.com/go-snart/snart/route"
)

// Plug is a pre-allocated Admin, to follow the same pattern as when loading with plugin.
var Plug = plug.Plug(&Admin{})

// Admin is a Plug with basic administration features.
type Admin struct {
	plug.Base
}

func (a *Admin) String() string {
	return "admin"
}

// PlugHandler adds the Admin's routes to the given Handler.
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
