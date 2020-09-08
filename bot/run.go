package bot

import (
	"context"
	"errors"
	"fmt"
	"time"

	dg "github.com/bwmarrin/discordgo"

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

	b.Session = b.Open(ctx)

	for _, p := range plugs {
		p.PlugSession(b.Session)
	}

	b.Session.AddHandler(b.Handler.Handle(b.DB))
	defer b.Session.Close()

	go b.cycleGamers()

	err := error(<-b.Halt)
	if err != nil {
		err = fmt.Errorf("halt: %w", err)
		log.Warn.Println(err)
		return err
	}

	return nil
}

func (b *Bot) cycleGamers() {
	for {
		game := b.Gamers.Random()()
		log.Debug.Printf("%v\n", game)

		for {
			err := b.Session.UpdateStatusComplex(
				dg.UpdateStatusData{Game: game},
			)
			if err == nil {
				break
			}

			if !errors.Is(err, dg.ErrWSNotFound) {
				err = fmt.Errorf("update status: %w", err)
				log.Warn.Println(err)

				break
			}

			time.Sleep(time.Second / 10)
		}

		time.Sleep(time.Second * 12)
	}
}
