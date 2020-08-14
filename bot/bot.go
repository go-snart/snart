// Package bot contains the general workings of a Snart Bot.
package bot

import (
	"time"

	dg "github.com/bwmarrin/discordgo"

	"github.com/go-snart/snart/db"
	"github.com/go-snart/snart/route"
)

// Bot holds all the internal workings of a Snart bot.
type Bot struct {
	Name string

	DB      *db.DB
	Session *dg.Session

	Handler *route.Handler
	Gamers  []Gamer

	Interrupt chan Interrupt
	Startup   time.Time

	Ready bool
}

// New creates a Bot.
func New() *Bot {
	return NewFromDB(db.New("bot"))
}

// NewFromDB creates a Bot from the given *db.DB.
func NewFromDB(d *db.DB) *Bot {
	return &Bot{
		DB:      d,
		Session: nil,

		Handler: &route.Handler{},
		Gamers:  []Gamer{GamerUptime},

		Interrupt: make(chan Interrupt),
		Startup:   time.Now(),

		Ready: false,
	}
}
