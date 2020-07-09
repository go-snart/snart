// Package bot contains the general workings of a Snart Bot.
package bot

import (
	"time"

	"github.com/go-snart/snart/db"
	"github.com/go-snart/snart/route"

	dg "github.com/bwmarrin/discordgo"
	"github.com/superloach/minori"
)

// Log is the logger for the bot package.
var Log = minori.GetLogger("bot")

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
	_f := "NewBot"
	b := &Bot{}

	Log.Debug(_f, "making bot")

	b.DB = d

	s, _ := dg.New()
	b.Session = s
	b.Session.AddHandler(b.Route)

	b.Router = route.NewRouter()
	b.Gamers = []Gamer{
		GamerUptime,
	}

	b.Interrupt = make(chan Interrupt)

	Log.Debug(_f, "made bot")

	return b, nil
}
