package token

import (
	"fmt"

	dg "github.com/bwmarrin/discordgo"

	"github.com/go-snart/snart/db"
	"github.com/go-snart/snart/log"
)

// Open opens a *dg.Session for you, pulling tokens from various sources.
func Open(d *db.DB, intents dg.Intent) *dg.Session {
	log.Debug.Println("enter->toks")

	toks := Tokens(d)

	log.Debug.Println("toks->tries")

	for _, tok := range toks {
		log.Debug.Println("tries->new")

		session, err := dg.New(tok)
		if err != nil {
			err = fmt.Errorf("new session %q: %w", tok, err)
			log.Warn.Println(err)
			log.Debug.Println("new->tries")

			continue
		}

		session.LogLevel = logLevel

		ready := make(chan *dg.Session)

		session.AddHandler(func(ses *dg.Session, _ *dg.Ready) {
			log.Info.Println("ready")
			ready <- ses
		})

		log.Debug.Println("new->open")

		session.Identify.Intents = dg.MakeIntent(intents)

		err = session.Open()
		if err != nil {
			err = fmt.Errorf("open %q: %w", tok, err)
			log.Warn.Println(err)
			log.Debug.Println("open->tries")

			continue
		}

		log.Debug.Printf("open %q->success", tok)

		return <-ready
	}

	log.Debug.Println("tries->exit")

	log.Info.Fatal("no suitable tokens found")

	return nil
}
