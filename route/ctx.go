package route

import dg "github.com/bwmarrin/discordgo"

type Ctx struct {
	Prefix      string
	CleanPrefix string
	Session     *dg.Session
	Message     *dg.Message
	Flags       *Flags
	Route       *Route
}

func (c *Ctx) Run() error {
	return c.Route.Func(c)
}
