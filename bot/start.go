package bot

import (
	"time"

	"github.com/go-snart/snart/db/token"
	"github.com/go-snart/snart/logs"
)

// Start performs the Bot's startup functions, and waits until an interrupt.
func (b *Bot) Start() {
	b.goPlugins()

	b.Session = token.Open(b.DB, &b.Ready)
	defer b.Session.Close()

	b.Session.AddHandler(b.Router.Handler(b.DB))

	b.Startup = time.Now()

	go b.cycleGamers()

	b.handleInterrupts()

	logs.Info.Println("bye :)")
}
