package bot

import (
	"context"
	"fmt"
	"time"

	"github.com/go-snart/snart/db/token"
)

// Start performs the Bot's startup functions, and then waits until an interrupt.
func (b *Bot) Start(ctx context.Context) error {
	const _f = "(*Bot).Start"

	b.GoPlugins()

	b.Startup = time.Now()

	tok, err := token.Token(ctx, b.DB)
	if err != nil {
		err = fmt.Errorf("token: %w", err)
		Log.Error(_f, err)

		return err
	}

	b.Session.Token = tok

	err = b.Session.Open()
	if err != nil {
		err = fmt.Errorf("session open: %w", err)
		Log.Error(_f, err)

		return err
	}

	go b.CycleGamers()

	b.WaitReady()
	Log.Info(_f, "ready")

	b.HandleInterrupts()

	b.Logout()

	return nil
}

// Logout performs standard disconnect routines.
func (b *Bot) Logout() {
	const _f = "(*Bot).Logout"

	err := b.Session.Close()
	if err != nil {
		err = fmt.Errorf("session close: %w", err)
		Log.Warn(_f, err)
	}

	Log.Info(_f, "logged out")
}
