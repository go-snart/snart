package snart

import (
	"context"
	"fmt"

	"github.com/go-snart/plug"
	"github.com/go-snart/snart/errs"
	"github.com/go-snart/snart/logs"
)

func (s *Snart) Load(name string) error {
	p, err := plug.Get(name)
	if err != nil {
		return fmt.Errorf("get plug %q: %w", err)
	}

	s.Plugs[name] = p

	p.PlugDB(s.DB)
	p.PlugHandler(s.Handler)
	p.PlugErr(s.Err)

	s.Intents |= p.Intents()
	b.Gamers = append(s.Gamers, p.Gamers()...)

	return nil
}

// Run performs the Bot's startup functions, and waits for an error.
func (s *Snart) Run(ctx context.Context) error {
	errs.Notify(b.Err)

	names := b.PlugNames()

	for _, name := range names {
		err := s.Load(name)
		if err != nil {
			err = fmt.Errorf("load %q: %w", err)
			logs.Warn.Println(err)

			continue
		}
	}

	defer b.State.Close()

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
