package test

import (
	"github.com/go-snart/snart/logs"
	"github.com/go-snart/snart/route"
)

var CtxBadSession = SessionBad()

func CtxBad() *route.Ctx {
	logs.Info.Println("enter")
	defer logs.Info.Println("exit")

	ctx := Ctx("")

	ctx.Session = CtxBadSession

	return ctx
}
