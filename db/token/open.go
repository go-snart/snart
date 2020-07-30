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

	Debug.Println("enter->toks")

	toks := Tokens(ctx, d)

	Debug.Println("toks->tries")

	for _, tok := range toks {
		Debug.Println("tries->new")

		session, _ := dg.New()
		session.Identify.Token = tok

		Debug.Println("new->open")

		err := session.Open()
		if err != nil {
			err = fmt.Errorf("open %q: %w", tok, err)
			Warn.Println(err)
		} else {
			Debug.Println("open->exit")
			opened = session
			return session
		}

		Debug.Println("open->tries")
	}

	Debug.Println("tries->exit")

	Info.Fatal("no suitable tokens found")

	return nil
}
