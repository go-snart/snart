package test

import (
	"context"
	"fmt"

	dg "github.com/bwmarrin/discordgo"
)

// SessionBot is a cached *bot.Bot for Session.
var SessionBot = Bot()

// Session gets a test *dg.Session.
func Session(ctx context.Context) (*dg.Session, error) {
	ses, err := SessionBot.Open(ctx)
	if err != nil {
		return nil, fmt.Errorf("test ses open: %w", err)
	}

	return ses, nil
}
