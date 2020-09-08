package test

import (
	"context"

	"github.com/go-snart/snart/bot"
)

// BotDB is a cached *db.DB for Bot.
var BotDB = DB(context.Background())

// Bot gets a test *bot.Bot.
func Bot() *bot.Bot {
	return bot.NewFromDB(BotDB)
}
