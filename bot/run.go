package bot

import (
	"errors"
	"fmt"
	"time"

	dg "github.com/bwmarrin/discordgo"

	"github.com/go-snart/snart/admin"
	"github.com/go-snart/snart/bot/plug"
	"github.com/go-snart/snart/db/token"
	"github.com/go-snart/snart/help"
	"github.com/go-snart/snart/log"
	logs "github.com/go-snart/snart/log"
)

// Run performs the Bot's startup functions, and waits for a Halt.
func (b *Bot) Run() {
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

	b.Session = token.Open(b.DB, b.Intents)

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
	}
}

func (b *Bot) cycleGamers() {
	for {
		for _, gamer := range b.Gamers {
			for {
				err := b.Session.UpdateStatusComplex(
					dg.UpdateStatusData{Game: gamer()},
				)
				if err == nil {
					break
				}

				if !errors.Is(err, dg.ErrWSNotFound) {
					err = fmt.Errorf("update status: %w", err)
					logs.Warn.Println(err)

					break
				}

				time.Sleep(time.Second / 10)
			}

			time.Sleep(time.Second * 12)
		}
	}
}
