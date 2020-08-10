package bot

import (
	"errors"
	"fmt"
	"time"

	dg "github.com/bwmarrin/discordgo"
	"github.com/go-snart/snart/logs"
)

// Gamer provides a Discord Game status for a given Bot state.
type Gamer func(*Bot) (*dg.Game, error)

func (b *Bot) cycleGamers() {
	b.WaitReady()

	for {
		for _, gamer := range b.Gamers {
			game, err := gamer(b)
			if err != nil {
				err = fmt.Errorf("gamer: %w", err)

				logs.Warn.Println(err)

				continue
			}

			for {
				err = b.Session.UpdateStatusComplex(
					dg.UpdateStatusData{
						Game: game,
					},
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
