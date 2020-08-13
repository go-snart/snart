package test

import (
	"github.com/go-snart/snart/logs"
	"github.com/go-snart/snart/route"
)

const (
	RouteName = "route"
	RouteCat  = "test"
	RouteDesc = "a test route"
)

var (
	RouteOkay  = route.True
	RouteMatch = route.MustMatch(`route|yeet`)
)

func RouteFunc(c *route.Ctx) error {
	logs.Info.Println("enter")
	defer logs.Info.Println("exit")

	c.Route.Desc = "run"
	return nil
}

func Route() *route.Route {
	logs.Info.Println("enter")
	defer logs.Info.Println("exit")

	return &route.Route{
		Name:  RouteName,
		Match: RouteMatch,
		Cat:   RouteCat,
		Desc:  RouteDesc,
		Okay:  RouteOkay,
		Func:  RouteFunc,
	}
}
