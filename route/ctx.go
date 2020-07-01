package route

import dg "github.com/bwmarrin/discordgo"

// Ctx holds a command context.
type Ctx struct {
	Prefix      string
	CleanPrefix string
	Session     *dg.Session
	Message     *dg.Message
	Flags       *Flags
	Route       *Route
}

// Run is a shortcut to c.Route.Func(c).
func (c *Ctx) Run() error {
	return c.Route.Func(c)
}
