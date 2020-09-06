package route

import (
	"fmt"

	dg "github.com/bwmarrin/discordgo"

	"github.com/go-snart/snart/log"
)

// Reply wraps a message to be sent to a given ChannelID using a given Session.
type Reply struct {
	*dg.MessageSend

	Session   *dg.Session
	ChannelID string
}

// Reply gets a Reply for the Ctx.
func (c *Ctx) Reply() *Reply {
	return &Reply{
		MessageSend: &dg.MessageSend{},
		Session:     c.Session,
		ChannelID:   c.Message.ChannelID,
	}
}

// SendMsg sends the Reply.
func (r *Reply) SendMsg() (*dg.Message, error) {
	return r.Session.ChannelMessageSendComplex(r.ChannelID, r.MessageSend)
}

// Send is a shortcut for SendMsg that log a warning on error and elides the resulting *dg.Message.
func (r *Reply) Send() error {
	_, err := r.SendMsg()
	if err != nil {
		err = fmt.Errorf("send msg: %w", err)
		log.Warn.Println(err)

		return err
	}

	return nil
}
