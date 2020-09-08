package test

import (
	"context"
	"strings"

	"github.com/go-snart/snart/route"
)

var (
	// CtxHandler is a cached *route.Handler for Ctx.
	CtxHandler = Handler()

	// CtxPrefix is a cached *prefix.Prefix for Ctx.
	CtxPrefix = Prefix()

	// CtxSession is a cached *dg.Session for Ctx.
	CtxSession = Session(context.Background())
)

// Ctx gets a test *route.Ctx.
func Ctx(ctx context.Context, content string) *route.Ctx {
	return CtxHandler.Ctx(
		ctx,
		CtxPrefix,
		CtxSession,
		Message(content),
		strings.Split(content, "\n")[0],
	)
}
