package test

import "github.com/go-snart/snart/route"

// HandlerRoute is a cached *route.Route for Handler.
var HandlerRoute = Route()

// Handler gets a test route.Handler.
func Handler() route.Handler {
	return route.Handler{
		HandlerRoute,
	}
}
