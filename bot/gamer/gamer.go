// Package gamer provides an abstraction for generating dynamic statuses.
package gamer

import dg "github.com/bwmarrin/discordgo"

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
