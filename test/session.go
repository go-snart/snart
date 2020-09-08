package test

import (
	"context"

	dg "github.com/bwmarrin/discordgo"
)

// SessionBot is a cached *bot.Bot for Session.
var SessionBot = Bot()

// Session gets a test *dg.Session.
func Session(ctx context.Context) *dg.Session {
	return SessionBot.Open(ctx)
}
