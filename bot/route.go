package bot

import (
	"fmt"
	"strings"

	"github.com/go-snart/snart/db"

	dg "github.com/bwmarrin/discordgo"
)

// Route is an event handler for dispatching a *dg.MessageCreate to the Bot's Router.
func (b *Bot) Route(s *dg.Session, m *dg.MessageCreate) {
	_f := "(*Bot).Route"
	Log.Debug(_f, "handling")

	if m.Message.Author.ID == s.State.User.ID {
		Log.Debug(_f, "ignore self")
		return
	}

	lines := strings.Split(m.Message.Content, "\n")
	Log.Debugf(_f, "lines %#v", lines)
	for _, line := range lines {
		Log.Debugf(_f, "line %#v", line)

		pfx, err := b.DB.FindPrefix(b.Session, m.GuildID, line)
		if err != nil {
			if err == db.ErrPrefixFail {
				continue
			}
			err = fmt.Errorf("prefix %#v %#v: %w", m.GuildID, line, err)
			Log.Warn(_f, err)
			continue
		}

		cpfx := pfx.Value
		if pfx.Clean != "" {
			cpfx = pfx.Clean
		}

		ctx := b.Router.Ctx(pfx.Value, cpfx, s, m.Message, line)
		if ctx == nil {
			Log.Warn(_f, "nil ctx")
			continue
		}

		err = ctx.Run()
		if err != nil {
			err = fmt.Errorf("ctx run: %w", err)
			Log.Warn(_f, err)
			continue
		}
	}
}
