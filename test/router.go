package test

import (
	"github.com/go-snart/snart/logs"
	"github.com/go-snart/snart/route"
)

var RouterRoute = Route()

func Router() *route.Router {
	logs.Info.Println("enter")
	defer logs.Info.Println("exit")

	router := route.NewRouter()

	router.Add(RouterRoute)

	return router
}
