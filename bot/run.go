package bot

import (
	"context"
	"fmt"

	"github.com/go-snart/snart/admin"
	"github.com/go-snart/snart/bot/halt"
	"github.com/go-snart/snart/bot/plug"
	"github.com/go-snart/snart/help"
	"github.com/go-snart/snart/log"
)

// Run performs the Bot's startup functions, and waits for a Halt.
func (b *Bot) Run(ctx context.Context) error {
	b.Halt = halt.Chan(ctx)

	plugs := append(
		[]plug.Plug{
			admin.Plug,
			help.Plug,
		},
		plug.Plugs(b.DB.Name)...,
	)

	for _, p := range plugs {
		p.PlugDB(b.DB)
		p.PlugHandler(b.Handler)
		p.PlugHalt(b.Halt)

		b.Intents |= p.Intents()
		b.Gamers = append(b.Gamers, p.Gamers()...)
	}

	ses, err := b.Open(ctx)
	if err != nil {
		err = fmt.Errorf("open ses: %w", err)
		log.Warn.Println(err)

		return err
	}
	defer ses.Close()

	b.Session = ses

	b.Session.AddHandler(b.Handler.Handle)

	for _, p := range plugs {
		p.PlugSession(b.Session)
	}

	go b.Gamers.Cycle(ctx, b.Session)

	err = error(<-b.Halt)
	if err != nil {
		err = fmt.Errorf("halt: %w", err)
		log.Warn.Println(err)

		return err
	}

	return nil
}
