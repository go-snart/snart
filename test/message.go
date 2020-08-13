package test

import (
	dg "github.com/bwmarrin/discordgo"
	"github.com/go-snart/snart/logs"
)

const (
	MessageID        = "12345678900"
	MessageChannelID = "12345678901"
	MessageGuildID   = "12345678902"
)

var MessageAuthor = &dg.User{}

func Message(content string) *dg.Message {
	logs.Info.Println("enter")
	defer logs.Info.Println("exit")

	return &dg.Message{
		ID:        MessageID,
		ChannelID: MessageChannelID,
		GuildID:   MessageGuildID,
		Content:   content,
		Author:    MessageAuthor,
	}
}
