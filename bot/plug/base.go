package plug

import (
	dg "github.com/bwmarrin/discordgo"

	"github.com/go-snart/snart/bot/gamer"
	"github.com/go-snart/snart/bot/halt"
	"github.com/go-snart/snart/db"
	"github.com/go-snart/snart/route"
)

type Base struct {
	DB      *db.DB
	Session *dg.Session
	Handler *route.Handler
	Halt    chan halt.Halt
}

func (b *Base) PlugDB(d *db.DB) {
	b.DB = d
}

func (b *Base) PlugSession(ses *dg.Session) {
	b.Session = ses
}

func (b *Base) PlugHandler(h *route.Handler) {
	b.Handler = h
}

func (b *Base) PlugHalt(h chan halt.Halt) {
	b.Halt = h
}

func (b *Base) Intents() dg.Intent {
	return 0
}

func (b *Base) Gamers() []gamer.Gamer {
	return nil
}
