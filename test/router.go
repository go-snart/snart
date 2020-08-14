package test

import "github.com/go-snart/snart/route"

// RouterRoute is a cached *route.Route for Router.
var RouterRoute = Route()

// Router gets a test *route.Router.
func Router() *route.Router {
	router := route.NewRouter()

	router.Add(RouterRoute)

	return router
}
