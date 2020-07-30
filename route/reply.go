package route

import dg "github.com/bwmarrin/discordgo"

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

// Send is a shortcut for SendMsg that logs a warning on error and elides the resulting *dg.Message.
func (r *Reply) Send() error {
	const _f = "(*Reply).Send"

	_, err := r.SendMsg()
	if err != nil {
		Log.Warn(_f, err)
	}

	return err
}
