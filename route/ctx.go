package route

import (
	dg "github.com/bwmarrin/discordgo"

	"github.com/go-snart/snart/db/prefix"
)

// Ctx holds a command context.
type Ctx struct {
	Prefix  *prefix.Prefix
	Session *dg.Session
	Message *dg.Message
	Flag    *Flag
	Route   *Route
}

// Run is a shortcut to c.Route.Func(c).
func (c *Ctx) Run() error {
	return c.Route.Func(c)
}
