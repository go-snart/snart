// Package bot contains the general workings of a Snart Bot.
package bot

import (
	"time"

	dg "github.com/bwmarrin/discordgo"

	"github.com/go-snart/snart/db"
	"github.com/go-snart/snart/logs"
	"github.com/go-snart/snart/route"
)

const _p = "bot"

var debug, info, warn = logs.Loggers(_p)

// Bot holds all the internal workings of a Snart bot.
type Bot struct {
	DB      *db.DB
	Session *dg.Session

	Router *route.Router
	Gamers []Gamer

	Interrupt chan Interrupt
	Startup   time.Time

	Ready chan struct{}
}

// New creates a Bot.
func New() *Bot {
	return &Bot{
		DB:      db.New(),
		Session: nil,

		Router: route.NewRouter(),
		Gamers: []Gamer{GamerUptime},

		Interrupt: make(chan Interrupt),
		Startup:   time.Now(),

		Ready: make(chan struct{}),
	}
}
