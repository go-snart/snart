package route

import dg "github.com/bwmarrin/discordgo"

// Wait contains two Okays, a Return channel, and a handler cancel function.
type Wait struct {
	General  Okay
	Specific Okay
	Return   chan *Ctx
	Cancel   func()
}

// WaitCancel creates a Wait from the Ctx and two Okays, with optional cancellation.
func (c *Ctx) WaitCancel(general Okay, specific Okay, cancel bool) *Wait {
	w := &Wait{}

	w.General = general
	w.Specific = specific
	w.Return = make(chan *Ctx)

	f := c.Session.AddHandler(w.Handle)

	if cancel {
		w.Cancel = f
	}

	return w
}

// Wait creates a Wait from the Ctx and two Okays, with cancellation enabled.
func (c *Ctx) Wait(general Okay, specific Okay) *Wait {
	return c.WaitCancel(general, specific, true)
}

func (w *Wait) Handle(s *dg.Session, m *dg.MessageCreate) {
	ctx := &Ctx{
		Session: s,
		Message: m.Message,
	}

	if !w.General(ctx) {
		return
	}

	if !w.Specific(ctx) {
		w.Return <- nil
	}

	w.Return <- ctx

	if w.Cancel != nil {
		w.Cancel()
	}
}
