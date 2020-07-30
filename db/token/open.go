package token

import (
	"context"
	"fmt"

	dg "github.com/bwmarrin/discordgo"

	"github.com/go-snart/snart/db"
)

var opened = (*dg.Session)(nil)

// Open opens a *dg.Session for you, pulling tokens from various sources.
func Open(ctx context.Context, d *db.DB) *dg.Session {
	const _f = "Open"

	if opened != nil {
		return opened
	}

	Log.Debug(_f, "enter->toks")

	toks := Tokens(ctx, d)

	Log.Debug(_f, "toks->tries")

	for _, tok := range toks {
		Log.Debug(_f, "tries->new")

		session, _ := dg.New()
		session.Identify.Token = tok

		Log.Debug(_f, "new->open")

		err := session.Open()
		if err != nil {
			err = fmt.Errorf("open %q: %w", tok, err)
			Log.Warn(_f, err)
		} else {
			Log.Debug(_f, "open->exit")
			opened = session
			return session
		}

		Log.Debug(_f, "open->tries")
	}

	Log.Debug(_f, "tries->exit")

	Log.Fatal("no suitable tokens found")

	return nil
}
