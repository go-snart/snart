package bot

import (
	"fmt"
	"time"

	dg "github.com/bwmarrin/discordgo"
)

// Gamer provides a Discord Game status for a given Bot state.
type Gamer func(*Bot) (*dg.Game, error)

// CycleGamers cycles through displaying the Gamers registered on the Bot.
func (b *Bot) CycleGamers() {
	_f := "(*Bot).CycleGamers"

	b.WaitReady()

	for {
		gamers := b.Gamers

		for _, gamer := range gamers {
			game, err := gamer(b)
			if err != nil {
				Log.Warn(_f, err)
				continue
			}

			err = b.Session.UpdateStatusComplex(dg.UpdateStatusData{Game: game})
			if err != nil {
				Log.Warn(_f, err)
				continue
			}

			time.Sleep(time.Second * 12)
		}
	}
}

// AddGamer registers a Gamer into the Bot.
func (b *Bot) AddGamer(g Gamer) {
	b.Gamers = append(b.Gamers, g)
}

// GamerUptime is an example Gamer that shows the Bot's uptime.
func GamerUptime(b *Bot) (*dg.Game, error) {
	return &dg.Game{
		Name: fmt.Sprintf("for %s", time.Since(b.Startup).Round(time.Second)),
		Type: dg.GameTypeGame,
	}, nil
}

// GamerText return an example Gamer that shows the given text and game type.
func GamerText(text string, typ dg.GameType) Gamer {
	return func(b *Bot) (*dg.Game, error) {
		return &dg.Game{Name: text, Type: typ}, nil
	}
}
