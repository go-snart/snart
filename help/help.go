// Package help is a Snart plug which provides a basic help menu.
package help

import (
	"time"

	"github.com/go-snart/snart/bot/plug"
	"github.com/go-snart/snart/route"
)

var Plug = plug.Plug(&Help{
	Startup: time.Now(),
})

type Help struct {
	plug.Base

	Startup time.Time
}

func (h *Help) String() string {
	return "help"
}

func (h *Help) PlugHandler(r *route.Handler) {
	h.Base.PlugHandler(r)

	r.Add(
		&route.Route{
			Name:  "help",
			Match: route.MustMatch("help"),
			Desc:  "help menu",
			Cat:   h.String(),
			Okay:  nil,
			Func:  h.Menu,
		},
		&route.Route{
			Name:  "about",
			Match: route.MustMatch("about|info"),
			Desc:  "info and stats about the bot",
			Cat:   h.String(),
			Okay:  nil,
			Func:  h.About,
		},
	)
}