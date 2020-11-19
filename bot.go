// Package bot contains the general workings of a Snart Bot.
package bot

import (
	"context"
	"time"

	dg "github.com/bwmarrin/discordgo"
	"github.com/go-snart/db"
	"github.com/go-snart/plug/gamer"
	"github.com/go-snart/route"
)

// Bot holds all the internal workings of a Snart bot.
type Bot struct {
	DB      *db.DB
	Session *dg.Session
	Handler *route.Handler
	Err     chan error

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
		Handler: route.NewHandler(d),
		Err:     nil,

		Intents: dg.IntentsAllWithoutPrivileged,
		Gamers: gamer.Queue{
			gamer.Uptime(time.Now()),
		},
	}
}
