package plugin

import (
	dg "github.com/bwmarrin/discordgo"

	"github.com/go-snart/snart/bot/gamer"
	"github.com/go-snart/snart/db"
	"github.com/go-snart/snart/route"
)

// Base is a Plugin that provides stub functionality.
type Base struct{}

// Session satisfies Plugin.
func (b Base) Session(_ *dg.Session) {}

// DB satisfies Plugin.
func (b Base) DB(_ *db.DB) {}

// Intents satisfies Plugin.
func (b Base) Intents() dg.Intent {
	return 0
}

// Routes satisfies Plugin.
func (b Base) Routes() []*route.Route {
	return nil
}

// Gamers satisfies Plugin.
func (b Base) Gamers() []gamer.Gamer {
	return nil
}

// String satisfies Plugin.
func (b Base) String() string {
	return ""
}
