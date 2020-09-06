package bot

import (
	"errors"
	"fmt"
	"time"

	dg "github.com/bwmarrin/discordgo"

	"github.com/go-snart/snart/db/token"
	logs "github.com/go-snart/snart/log"
)

// Start performs the Bot's startup functions, and waits until an interrupt.
func (b *Bot) Start() {
	for _, p := range b.Plugins {
		p.DB(b.DB)

		b.Intents |= p.Intents()

		go p.Session(b.Session)

		b.Handler.Add(p.Routes()...)

		b.Gamers = append(b.Gamers, p.Gamers()...)
	}

	b.Session = token.Open(b.DB, b.Intents)
	defer b.Session.Close()

	b.Session.AddHandler(b.Handler.Handle(b.DB))

	b.Startup = time.Now()

	go b.cycleGamers()

	b.handleInterrupts()
	logs.Info.Println("bye :)")
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
