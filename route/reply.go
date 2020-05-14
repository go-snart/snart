package route

import dg "github.com/bwmarrin/discordgo"

type Reply struct {
	*dg.MessageSend

	Session   *dg.Session
	ChannelID string
}

func (r *Reply) SendMsg() (*dg.Message, error) {
	return r.Session.ChannelMessageSendComplex(r.ChannelID, r.MessageSend)
}

func (r *Reply) Send() error {
	_f := "(*Reply).Send"
	_, err := r.SendMsg()
	if err != nil {
		Log.Warn(_f, err)
	}
	return err
}

func (c *Ctx) Reply() *Reply {
	return &Reply{
		MessageSend: &dg.MessageSend{},
		Session:     c.Session,
		ChannelID:   c.Message.ChannelID,
	}
}
