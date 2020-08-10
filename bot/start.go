package bot

import (
	"context"
	"fmt"
	"time"

	"github.com/go-snart/snart/db/token"
	"github.com/go-snart/snart/logs"
)

// Start performs the Bot's startup functions, and then waits until an interrupt.
func (b *Bot) Start(ctx context.Context) {
	b.GoPlugins()

	b.Session = token.Open(ctx, b.DB)
	b.Session.AddHandler(b.Router.Handler(b.DB))
	b.Session.LogLevel = logs.DGLevel

	b.Startup = time.Now()

	go b.CycleGamers()

	b.WaitReady()
	info.Println("ready")

	b.HandleInterrupts()

	b.Logout()
}

// Logout performs standard disconnect routines.
func (b *Bot) Logout() {
	err := b.Session.Close()
	if err != nil {
		err = fmt.Errorf("session close: %w", err)
		warn.Println(err)
	}

	info.Println("logged out")
}
