// Package gamer provides an abstraction for generating dynamic statuses.
package gamer

import (
	"fmt"
	"time"

	dg "github.com/bwmarrin/discordgo"
)

// Gamer provides a Discord Game status.
type Gamer func() *dg.Game

// Text returns a Gamer with the given name and type.
func Text(name string, typ dg.GameType) Gamer {
	return func() *dg.Game {
		return &dg.Game{
			Name: name,
			Type: typ,
		}
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
