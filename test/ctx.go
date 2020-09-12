package test

import (
	"context"
	"strings"

	dg "github.com/bwmarrin/discordgo"

	"github.com/go-snart/snart/route"
)

var (
	// CtxHandler is a cached *route.Handler for Ctx.
	CtxHandler = Handler()

	// CtxPrefix is a cached *prefix.Prefix for Ctx.
	CtxPrefix = Prefix()

	// CtxSession is a cached *dg.Session for Ctx.
	CtxSession = func() *dg.Session {
		ses, err := Session(context.Background())
		if err != nil {
			panic(err)
		}

		return ses
	}()
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
