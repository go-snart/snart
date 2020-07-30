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
	if opened != nil {
		return opened
	}

	debug.Println("enter->toks")

	toks := Tokens(ctx, d)

	debug.Println("toks->tries")

	for _, tok := range toks {
		debug.Println("tries->new")

		session, _ := dg.New()
		session.Identify.Token = tok

		debug.Println("new->open")

		err := session.Open()
		if err != nil {
			err = fmt.Errorf("open %q: %w", tok, err)
			warn.Println(err)
		} else {
			debug.Println("open->exit")
			opened = session
			return session
		}

		debug.Println("open->tries")
	}

	debug.Println("tries->exit")

	info.Fatal("no suitable tokens found")

	return nil
}
