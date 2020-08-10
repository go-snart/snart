package bot

import (
	"context"
	"time"

	"github.com/go-snart/snart/db/token"
)

// Start performs the Bot's startup functions, and waits until an interrupt.
func (b *Bot) Start(ctx context.Context) {
	b.goPlugins()

	b.Session = token.Open(ctx, b.DB, &b.Ready)
	defer b.Session.Close()

	b.Session.AddHandler(b.Router.Handler(b.DB))

	b.Startup = time.Now()

	go b.cycleGamers()

	b.handleInterrupts()
}
