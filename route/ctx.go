package route

import (
	"context"

	dg "github.com/bwmarrin/discordgo"

	"github.com/go-snart/snart/db"
)

// Ctx holds a command context.
type Ctx struct {
	context.Context

	Prefix  *db.Prefix
	Session *dg.Session
	Message *dg.Message
	Flag    *Flag
	Route   *Route
}

// Run is a shortcut to c.Route.Func(c).
func (c *Ctx) Run() error {
	return c.Route.Func(c)
}
