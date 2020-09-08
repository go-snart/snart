// Package bot contains the general workings of a Snart Bot.
package bot

import (
	"context"

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
	Handler *route.Handler
	Halt    chan halt.Halt

	Intents dg.Intent
	Gamers  gamer.Queue
}

// New creates a Bot.
func New(ctx context.Context, name string) *Bot {
	return NewFromDB(db.New(ctx, name))
}

// NewFromDB creates a Bot from the given *db.DB.
func NewFromDB(d *db.DB) *Bot {
	return &Bot{
		DB:      d,
		Session: nil,
		Halt:    nil,

		Handler: &route.Handler{},
		Intents: dg.IntentsAllWithoutPrivileged,
		Gamers:  gamer.Queue{gamer.Uptime()},
	}
}
