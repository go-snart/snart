package route

import (
	"context"
	"time"

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
	ctx     context.Context
}

// Deadline satisfies context.Context.
func (c *Ctx) Deadline() (time.Time, bool) {
	return c.ctx.Deadline()
}

// Done satisfies context.Context.
func (c *Ctx) Done() <-chan struct{} {
	return c.ctx.Done()
}

// Err satisfies context.Context.
func (c *Ctx) Err() error {
	return c.ctx.Err()
}

// Value satisfies context.Context.
func (c *Ctx) Value(key interface{}) interface{} {
	return c.ctx.Value(key)
}

// Run is a shortcut to c.Route.Func(c).
func (c *Ctx) Run() error {
	return c.Route.Func(c)
}
