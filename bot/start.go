package bot

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-snart/snart/db/token"
)

// ErrAllToksFailed occurs when all of the provided tokens failed to authenticate.
var ErrAllToksFailed = errors.New("all tokens failed")

// Start performs the Bot's startup functions, and then waits until an interrupt.
func (b *Bot) Start(ctx context.Context) error {
	b.GoPlugins()

	b.Session = token.Open(ctx, b.DB)
	b.Session.AddHandler(b.Router.Handler(b.DB))

	b.Startup = time.Now()

	go b.CycleGamers()

	b.WaitReady()
	Info.Println("ready")

	b.HandleInterrupts()

	b.Logout()

	return nil
}

// Logout performs standard disconnect routines.
func (b *Bot) Logout() {
	err := b.Session.Close()
	if err != nil {
		err = fmt.Errorf("session close: %w", err)
		Warn.Println(err)
	}

	Info.Println("logged out")
}
