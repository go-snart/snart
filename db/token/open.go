package token

import (
	"fmt"

	dg "github.com/bwmarrin/discordgo"

	"github.com/go-snart/snart/db"
	"github.com/go-snart/snart/logs"
)

// Open opens a *dg.Session for you, pulling tokens from various sources.
func Open(d *db.DB, ready *bool) *dg.Session {
	logs.Debug.Println("enter->toks")

	toks := Tokens(d)

	logs.Debug.Println("toks->tries")

	for _, tok := range toks {
		logs.Debug.Println("tries->new")

		session, err := dg.New(tok)
		if err != nil {
			err = fmt.Errorf("new session %q: %w", tok, err)
			logs.Warn.Println(err)
			logs.Debug.Println("new->tries")

			continue
		}

		session.LogLevel = logLevel

		if ready != nil {
			session.AddHandler(func(_ *dg.Session, _ *dg.Ready) {
				*ready = true
				logs.Info.Println("ready")
			})
		}

		logs.Debug.Println("new->open")

		err = session.Open()
		if err != nil {
			err = fmt.Errorf("open %q: %w", tok, err)
			logs.Warn.Println(err)
			logs.Debug.Println("open->tries")

			continue
		}

		logs.Debug.Printf("open %q->success", tok)

		return session
	}

	logs.Debug.Println("tries->exit")

	logs.Info.Fatal("no suitable tokens found")

	return nil
}
