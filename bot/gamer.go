package bot

import (
	"fmt"
	"time"

	dg "github.com/bwmarrin/discordgo"
)

type Gamer func(*Bot) (*dg.Game, error)

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

func (b *Bot) AddGamer(g Gamer) {
	b.Gamers = append(b.Gamers, g)
}

func GamerUptime(b *Bot) (*dg.Game, error) {
	return &dg.Game{
		Name: fmt.Sprintf("for %s", time.Now().Sub(b.Startup).Round(time.Second)),
		Type: dg.GameTypeGame,
	}, nil
}
