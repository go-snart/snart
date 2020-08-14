package test

import (
	"strings"

	"github.com/go-snart/snart/route"
)

var (
	// CtxHandler is a cached route.Handler for Ctx.
	CtxHandler = Handler()

	// CtxPrefix is a cached *prefix.Prefix for Ctx.
	CtxPrefix = Prefix()

	// CtxSession is a cached *dg.Session for Ctx.
	CtxSession = Session()
)

// Ctx gets a test *route.Ctx.
func Ctx(content string) *route.Ctx {
	return CtxHandler.Ctx(
		CtxPrefix,
		CtxSession,
		Message(content),
		strings.Split(content, "\n")[0],
	)
}
