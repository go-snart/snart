package test

import dg "github.com/bwmarrin/discordgo"

// Intent gets a test dg.Intent.
func Intent() dg.Intent {
	return dg.IntentsAll
}
