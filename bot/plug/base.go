package plug

import (
	dg "github.com/bwmarrin/discordgo"

	"github.com/go-snart/snart/bot/halt"
	"github.com/go-snart/snart/db"
)

type Base struct {
	DB      *db.DB
	Session *dg.Session
	Halt    chan halt.Halt
}

func (b *Base) PlugDB(d *db.DB) {
	b.DB = d
}

func (b *Base) PlugSession(ses *dg.Session) {
	b.Session = ses
}

func (b *Base) PlugHalt(h chan halt.Halt) {
	b.Halt = h
}
