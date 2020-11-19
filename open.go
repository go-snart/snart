package bot

import (
	"context"
	"errors"
	"fmt"

	dg "github.com/bwmarrin/discordgo"
	"github.com/go-snart/logs"
)

// ErrNoSession occurs when Open is unable to open a session.
var ErrNoSession = errors.New("no session found")

// Open opens a *dg.Session for you, pulling tokens from various sources.
func (b *Bot) Open(ctx context.Context) (*dg.Session, error) {
	logs.Debug.Println("enter->toks")

	toks, err := b.DB.Tokens(ctx)
	if err != nil {
		err = fmt.Errorf("toks: %w", err)
		logs.Warn.Println(err)

		return nil, err
	}

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

		ready := make(chan *dg.Session)

		session.AddHandler(func(ses *dg.Session, _ *dg.Ready) {
			logs.Info.Println("ready")
			ready <- ses
		})

		logs.Debug.Println("new->open")

		session.Identify.Intents = dg.MakeIntent(b.Intents)

		err = session.Open()
		if err != nil {
			err = fmt.Errorf("open %q: %w", tok, err)
			logs.Warn.Println(err)
			logs.Debug.Println("open->tries")

			continue
		}

		logs.Debug.Printf("open %q->success", tok)

		return <-ready, nil
	}

	logs.Debug.Println("tries->exit")

	return nil, ErrNoSession
}
