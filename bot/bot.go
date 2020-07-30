// Package bot contains the general workings of a Snart Bot.
package bot

import (
	"fmt"
	"time"

	dg "github.com/bwmarrin/discordgo"

	"github.com/go-snart/snart/db"
	"github.com/go-snart/snart/route"
)

// Bot holds all the internal workings of a Snart bot.
type Bot struct {
	DB      *db.DB
	Session *dg.Session

	Router *route.Router
	Gamers []Gamer

	Interrupt chan Interrupt
	Startup   time.Time
}

// NewBot creates a Bot from a given DB connection.
func NewBot(d *db.DB) (*Bot, error) {
	const _f = "NewBot"

	Log.Debug(_f, "making bot")
	defer Log.Debug(_f, "made bot")

	ses, err := dg.New()
	if err != nil {
		err = fmt.Errorf("new ses: %w", err)
		Log.Error(_f, err)

		return nil, err
	}

	router := route.NewRouter()
	ses.AddHandler(router.Handler(d))

	return &Bot{
		DB:      d,
		Session: ses,

		Router: router,
		Gamers: []Gamer{GamerUptime},

		Interrupt: make(chan Interrupt),
		Startup:   time.Now(),
	}, nil
}
