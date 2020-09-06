package admin

import (
	"errors"
	"fmt"

	"github.com/go-snart/snart/bot/halt"
	"github.com/go-snart/snart/log"
	"github.com/go-snart/snart/route"
)

// ErrRestart occurs when a restart is requested.
var ErrRestart = errors.New("restart")

// Restart is a command that restarts the Bot.
func (a *Admin) Restart(ctx *route.Ctx) error {
	err := ctx.Flag.Parse()
	if err != nil {
		err = fmt.Errorf("flag parse: %w", err)
		log.Warn.Println(err)

		return err
	}

	log.Warn.Println("restarting...")

	rep := ctx.Reply()
	rep.Content = "restarting..."

	err = rep.Send()

	a.Halt <- halt.Halt{Err: ErrRestart}

	return err
}
