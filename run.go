package bot

import (
	"context"
	"fmt"

	"github.com/go-snart/logs"
	"github.com/go-snart/plug"
	"github.com/go-snart/plug/admin"
	"github.com/go-snart/plug/help"
)

// Run performs the Bot's startup functions, and waits for an error.
func (b *Bot) Run(ctx context.Context) error {
	b.Err = SigChan(ctx)

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
		p.PlugErr(b.Err)

		b.Intents |= p.Intents()
		b.Gamers = append(b.Gamers, p.Gamers()...)
	}

	ses, err := b.Open(ctx)
	if err != nil {
		err = fmt.Errorf("open ses: %w", err)
		logs.Warn.Println(err)

		return err
	}
	defer ses.Close()

	b.Session = ses

	b.Session.AddHandler(b.Handler.Handle)

	for _, p := range plugs {
		p.PlugSession(b.Session)
	}

	go b.Gamers.Cycle(ctx, b.Session)

	err = <-b.Err
	if err != nil {
		logs.Warn.Println(err)
		return err
	}

	return nil
}
