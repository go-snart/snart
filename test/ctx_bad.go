package test

import "github.com/go-snart/snart/route"

// CtxBadSession is a cached *dg.Session for CtxBad.
var CtxBadSession = SessionBad()

// CtxBad gets a bad test *route.Ctx.
func CtxBad() *route.Ctx {
	ctx := Ctx("")

	ctx.Session = CtxBadSession

	return ctx
}
