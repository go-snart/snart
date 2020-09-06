// Package bot contains the general workings of a Snart Bot.
package bot

import (
	"time"

	dg "github.com/bwmarrin/discordgo"

	"github.com/go-snart/snart/bot/gamer"
	"github.com/go-snart/snart/bot/plugin"
	"github.com/go-snart/snart/db"
	"github.com/go-snart/snart/route"
)

// Bot holds all the internal workings of a Snart bot.
type Bot struct {
	DB      *db.DB
	Session *dg.Session

	Plugins []plugin.Plugin

	Intents   dg.Intent
	Handler   *route.Handler
	Gamers    []gamer.Gamer
	Interrupt chan Interrupt
	Startup   time.Time
}

// New creates a Bot.
func New() *Bot {
	return NewFromDB(db.New("bot"))
}

// NewFromDB creates a Bot from the given *db.DB.
func NewFromDB(d *db.DB) *Bot {
	b := &Bot{
		DB:      d,
		Session: nil,

		Plugins: plugin.OpenAll(),

		Intents:   dg.IntentsAllWithoutPrivileged,
		Handler:   &route.Handler{},
		Gamers:    nil,
		Interrupt: make(chan Interrupt),
		Startup:   time.Now(),
	}

	b.Gamers = append(b.Gamers, b.Uptime)

	return b
}
