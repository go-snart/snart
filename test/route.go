package test

import "github.com/go-snart/snart/route"

const (
	// RouteName is the route name used by Route.
	RouteName = "route"

	// RouteCat is the route category used by Route.
	RouteCat = "test"

	// RouteDesc is the route description used by Route.
	RouteDesc = "a test route"

	// RouteDescNew is the updated route description used by RouteFunc.
	RouteDescNew = "run"
)

var (
	// RouteOkay is the route.Okay used by Route.
	RouteOkay = route.True

	// RouteMatch is the *re2.Regexp used by Route.
	RouteMatch = route.MustMatch(`route|yeet`)
)

// RouteFunc is the route func used by Route.
func RouteFunc(c *route.Ctx) error {
	c.Route.Desc = RouteDescNew
	return nil
}

// Route gets a test *route.Route.
func Route() *route.Route {
	return &route.Route{
		Name:  RouteName,
		Match: RouteMatch,
		Cat:   RouteCat,
		Desc:  RouteDesc,
		Okay:  RouteOkay,
		Func:  RouteFunc,
	}
}
