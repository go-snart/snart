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
	b.WaitReady()

	for {
		for _, gamer := range b.Gamers {
			game, err := gamer(b)
			if err != nil {
				err = fmt.Errorf("gamer: %w", err)

				warn.Println(err)

				continue
			}

			err = b.Session.UpdateStatusComplex(
				dg.UpdateStatusData{
					Game: game,
				},
			)
			if err != nil {
				err = fmt.Errorf("update status: %w", err)

				warn.Println(err)

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
func GamerText(name string, typ dg.GameType) Gamer {
	game := &dg.Game{
		Name: name,
		Type: typ,
	}

	return func(_ *Bot) (*dg.Game, error) {
		return game, nil
	}
}
