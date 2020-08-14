package test

import dg "github.com/bwmarrin/discordgo"

const (
	// MessageID is the message id used by Message.
	MessageID = "12345678902"

	// MessageChannelID is the channel id used by Message.
	MessageChannelID = "12345678901"

	// MessageGuildID is the guild id used by Message.
	MessageGuildID = "12345678900"
)

// MessageAuthor is the *dg.User used by Message.
var MessageAuthor = &dg.User{}

// Message gets a test *dg.Message with the given content.
func Message(content string) *dg.Message {
	return &dg.Message{
		ID:        MessageID,
		ChannelID: MessageChannelID,
		GuildID:   MessageGuildID,
		Content:   content,
		Author:    MessageAuthor,
	}
}
