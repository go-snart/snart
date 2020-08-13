package test

import (
	dg "github.com/bwmarrin/discordgo"

	"github.com/go-snart/snart/db/token"
	"github.com/go-snart/snart/logs"
)

var SessionDB = DB()

func Session() *dg.Session {
	logs.Info.Println("enter")
	defer logs.Info.Println("exit")

	return token.Open(SessionDB, nil)
}
