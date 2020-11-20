package snart

import (
	"context"
	"fmt"

	"github.com/go-snart/plug"
	"github.com/go-snart/snart/errs"
	"github.com/go-snart/snart/logs"
)

func (b *Bot) Load(name string) error {
	p, err := plug.Get(name)
	if err != nil {
		return fmt.Errorf("get plug %q: %w", err)
	}

	b.Plugs[name] = p

	p.PlugDB(b.DB)
	p.PlugHandler(b.Handler)
	p.PlugErr(b.Err)
	p.PlugState(b.Err)

	b.Intents |= p.Intents()
	b.Gamers = append(b.Gamers, p.Gamers()...)

	return nil
}

// Run performs the Bot's startup functions, and waits for an error.
func (b *Bot) Run(ctx context.Context) error {
	errs.Notify(b.Err)

	names := b.PlugNames()

	for _, name := range names {
		err := b.Load(name)
		if err != nil {
			err = fmt.Errorf("load %q: %w", err)
			logs.Warn.Println(err)

			continue
		}
	}

	defer b.State.Close()

	b.State.AddHandler(b.Route.Handle)

	go b.Gamers.Cycle(b.State)

	err = <-b.Err
	if err != nil {
		logs.Warn.Println(err)
		return err
	}

	return nil
}
