package route

import (
	"context"
	"time"

	dg "github.com/bwmarrin/discordgo"
)

// Ctx holds a command context.
type Ctx struct {
	Prefix      string
	CleanPrefix string
	Session     *dg.Session
	Message     *dg.Message
	Flags       *Flags
	Route       *Route
	Context     context.Context
}

func (c *Ctx) Deadline() (time.Time, bool) {
	return c.Context.Deadline()
}

func (c *Ctx) Done() <-chan struct{} {
	return c.Context.Done()
}

func (c *Ctx) Err() error {
	return c.Context.Err()
}

func (c *Ctx) Value(key interface{}) interface{} {
	return c.Context.Value(key)
}

// Run is a shortcut to c.Route.Func(c).
func (c *Ctx) Run() error {
	return c.Route.Func(c)
}
