package test

import (
	dg "github.com/bwmarrin/discordgo"
	"github.com/go-snart/snart/logs"
)

const SessionBadToken = "session bad token"

func SessionBad() *dg.Session {
	logs.Info.Println("enter")
	defer logs.Info.Println("exit")

	session, _ := dg.New(SessionBadToken)
	return session
}
