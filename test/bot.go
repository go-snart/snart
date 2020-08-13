package test

import (
	"github.com/go-snart/snart/bot"
	"github.com/go-snart/snart/logs"
)

var BotDB = DB()

func Bot() *bot.Bot {
	logs.Info.Println("enter")
	defer logs.Info.Println("exit")

	return bot.NewFromDB(BotDB)
}
