package plug

import (
	dg "github.com/bwmarrin/discordgo"

	"github.com/go-snart/snart/bot/gamer"
	"github.com/go-snart/snart/bot/halt"
	"github.com/go-snart/snart/db"
	"github.com/go-snart/snart/route"
)

// Base is a Plug that provides stub functionality to be embedded into other Plugs.
type Base struct {
	DB      *db.DB
	Session *dg.Session
	Handler *route.Handler
	Halt    chan halt.Halt
}

// PlugDB loads the given db.DB into the Base.
func (b *Base) PlugDB(d *db.DB) {
	b.DB = d
}

// PlugSession loads the given dg.Session into the Base.
func (b *Base) PlugSession(ses *dg.Session) {
	b.Session = ses
}

// PlugHandler loads the given route.Handler into the Base.
func (b *Base) PlugHandler(h *route.Handler) {
	b.Handler = h
}

// PlugHalt loads the given halt.Halt channel into the Base.
func (b *Base) PlugHalt(h chan halt.Halt) {
	b.Halt = h
}

// Intents returns additional dg.Intents to be ORed in.
func (b *Base) Intents() dg.Intent {
	return 0
}

// Gamers returns additional gamer.Gamers to be added to the Queue.
func (b *Base) Gamers() []gamer.Gamer {
	return nil
}
