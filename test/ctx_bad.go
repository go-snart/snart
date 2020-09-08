package test

import (
	"context"

	"github.com/go-snart/snart/route"
)

// CtxBadSession is a cached *dg.Session for CtxBad.
var CtxBadSession = SessionBad()

// CtxBad gets a bad test *route.Ctx.
func CtxBad(ctx context.Context) *route.Ctx {
	c := Ctx(ctx, "")

	c.Session = CtxBadSession

	return c
}
