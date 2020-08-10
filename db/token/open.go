package token

import (
	"context"
	"fmt"

	dg "github.com/bwmarrin/discordgo"

	"github.com/go-snart/snart/db"
)

// Open opens a *dg.Session for you, pulling tokens from various sources.
func Open(ctx context.Context, d *db.DB) *dg.Session {
	debug.Println("enter->toks")

	toks := Tokens(ctx, d)

	debug.Println("toks->tries")

	for _, tok := range toks {
		debug.Println("tries->new")

		session, err := dg.New(tok)
		if err != nil {
			err = fmt.Errorf("new session %q: %w", tok, err)
			warn.Println(err)
			debug.Println("new->tries")

			continue
		}

		debug.Println("new->open")

		err = session.Open()
		if err != nil {
			err = fmt.Errorf("open %q: %w", tok, err)
			warn.Println(err)
			debug.Println("open->tries")

			continue
		}

		debug.Printf("open %q->success", tok)

		return session
	}

	debug.Println("tries->exit")

	info.Fatal("no suitable tokens found")

	return nil
}
