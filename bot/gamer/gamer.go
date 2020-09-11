// Package gamer provides Discord Game status generation.
package gamer

import (
	"fmt"
	"time"

	dg "github.com/bwmarrin/discordgo"
)

// Gamer is a closure that generates a Discord Game status.
type Gamer interface {
	Game() *dg.Game
}

// Static is a Gamer that returns itself.
type Static dg.Game

// Game implements Gamer.
func (s *Static) Game() *dg.Game {
	return (*dg.Game)(s)
}

// Uptime is a Gamer that shows the duration since its value.
type Uptime time.Time

// Game implements Gamer.
func (u Uptime) Game() *dg.Game {
	return &dg.Game{
		Name: fmt.Sprintf("for %s", time.Since(time.Time(u)).Round(time.Second)),
		Type: dg.GameTypeGame, // playing...
	}
}
