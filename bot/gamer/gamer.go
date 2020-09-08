// Package gamer provides Discord Game status generation.
package gamer

import (
	"fmt"
	"time"

	dg "github.com/bwmarrin/discordgo"
)

// Gamer is a closure that generates a Discord Game status.
type Gamer func() *dg.Game

// Text returns a Gamer with the given name and type.
func Text(name string, typ dg.GameType) Gamer {
	game := &dg.Game{
		Name: name,
		Type: typ,
	}

	return func() *dg.Game {
		return game
	}
}

// Uptime returns a Gamer that shows the uptime (since Uptime was called).
func Uptime() Gamer {
	startup := time.Now()

	return func() *dg.Game {
		return &dg.Game{
			Name: fmt.Sprintf("for %s", time.Since(startup).Round(time.Second)),
			Type: dg.GameTypeGame, // playing...
		}
	}
}
