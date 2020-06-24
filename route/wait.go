package route

import dg "github.com/bwmarrin/discordgo"

type Wait struct {
	general  Okay
	specific Okay
	Return   chan *Ctx
	cancel   *func()
}

func (c *Ctx) WaitCancel(general Okay, specific Okay, cancel bool) *Wait {
	w := &Wait{}

	w.general = general
	w.specific = specific
	w.Return = make(chan *Ctx)

	f := c.Session.AddHandler(w.handle)
	if cancel {
		w.cancel = &f
	} else {
		w.cancel = nil
	}

	return w
}

func (c *Ctx) Wait(general Okay, specific Okay) *Wait {
	return c.WaitCancel(general, specific, true)
}

func (w *Wait) handle(s *dg.Session, m *dg.MessageCreate) {
	ctx := &Ctx{
		Session: s,
		Message: m.Message,
	}

	if !w.general(ctx) {
		return
	}

	if !w.specific(ctx) {
		w.Return <- nil
	}

	w.Return <- ctx

	if w.cancel != nil {
		(*w.cancel)()
	}
}
