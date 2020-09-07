package bot

import (
	"context"
	"fmt"

	dg "github.com/bwmarrin/discordgo"

	"github.com/go-snart/snart/log"
)

// Open opens a *dg.Session for you, pulling tokens from various sources.
func (b *Bot) Open(ctx context.Context) *dg.Session {
	log.Debug.Println("enter->toks")

	toks := b.DB.Tokens(ctx)

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

		ready := make(chan *dg.Session)

		session.AddHandler(func(ses *dg.Session, _ *dg.Ready) {
			log.Info.Println("ready")
			ready <- ses
		})

		log.Debug.Println("new->open")

		session.Identify.Intents = dg.MakeIntent(b.Intents)

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
