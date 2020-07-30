// Package bot contains the general workings of a Snart Bot.
package bot

import (
	"context"
	"fmt"
	"time"

	dg "github.com/bwmarrin/discordgo"

	"github.com/go-snart/snart/db"
	"github.com/go-snart/snart/db/token"
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

// New creates a Bot.
func New(ctx context.Context) (*Bot, error) {
	const _f = "New"

	Log.Debug(_f, "making bot")
	defer Log.Debug(_f, "made bot")

	d, err := db.New()
	if err != nil {
		err = fmt.Errorf("new db: %w", err)
		Log.Error(_f, err)

		return nil, err
	}

	ses := token.Open(ctx, d)

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
