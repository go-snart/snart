// Package bot contains the general workings of a Snart Bot.
package bot

import (
	dg "github.com/bwmarrin/discordgo"

	"github.com/go-snart/snart/bot/gamer"
	"github.com/go-snart/snart/bot/halt"
	"github.com/go-snart/snart/db"
	"github.com/go-snart/snart/route"
)

// Bot holds all the internal workings of a Snart bot.
type Bot struct {
	DB      *db.DB
	Session *dg.Session
	Halt    chan halt.Halt

	Handler *route.Handler
	Intents dg.Intent
	Gamers  []gamer.Gamer
}

// New creates a Bot.
func New(name string) *Bot {
	return NewFromDB(db.New(name))
}

// NewFromDB creates a Bot from the given *db.DB.
func NewFromDB(d *db.DB) *Bot {
	return &Bot{
		DB:      d,
		Session: nil,
		Halt:    halt.Chan(),

		Handler: &route.Handler{},
		Intents: dg.IntentsAllWithoutPrivileged,
		Gamers:  []gamer.Gamer{gamer.Uptime()},
	}
}
