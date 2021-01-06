package admin

import (
	"errors"
	"log"

	"github.com/go-snart/route"
)

// ErrRestart occurs when a restart is requested.
var ErrRestart = errors.New("restart")

// Restart is a route.Func that kills the Bot.
func (a *Admin) Restart(t *route.Trigger) error {
	log.Println("restarting...")

	rep := t.Reply()
	rep.Content = "restarting..."
	_ = rep.Send()

	a.Bot.Errs <- ErrRestart

	return nil
}
