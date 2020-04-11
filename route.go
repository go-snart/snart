package bot

import (
	"strings"

	"github.com/go-snart/snart/lib/errs"
	"github.com/go-snart/snart/lib/route"

	dg "github.com/bwmarrin/discordgo"
)

func (b *Bot) AddRoute(rs ...*route.Route) {
	b.Routes = append(b.Routes, rs...)
}

func (b *Bot) Route(s *dg.Session, m *dg.MessageCreate) {
	_f := "(*Bot).Route"
	Log.Debug(_f, "handling")

	if m.Message.Author.ID == s.State.User.ID {
		Log.Debug(_f, "ignore self")
		return
	}

	for _, line := range strings.Split(m.Message.Content, "\n") {
		msg := &(*m.Message)
		msg.Content = line

		pfx, cpfx, err := b.Prefix(msg.GuildID, msg.Content)
		if err != nil {
			if err == PrefixFail {
				continue
			}
			errs.Wrap(&err, `b.Prefix(%#v, %#v)`, msg.GuildID, msg.Content)
			Log.Warn(_f, err)
			continue
		}

		ctx, err := route.GetCtx(pfx, cpfx, s, msg, b.Routes)
		if err != nil {
			errs.Wrap(&err, `route.GetCtx(GetPrefix(b.DB, s), s, msg, b.Routes)`)
			Log.Warn(_f, err)
			continue
		}

		if ctx == nil {
			continue
		}

		err = ctx.Run()
		if err != nil {
			errs.Wrap(&err, `ctx.Run()`)
			Log.Warn(_f, err)
			continue
		}
	}
}
