package test

import (
	dg "github.com/bwmarrin/discordgo"
	"github.com/go-snart/snart/logs"
)

func MessageCreate(content string) *dg.MessageCreate {
	logs.Info.Println("enter")
	defer logs.Info.Println("exit")

	return &dg.MessageCreate{
		Message: Message(content),
	}
}
