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
	const _f = "(*Bot).Start"

	b.GoPlugins()

	b.Startup = time.Now()

	toks := token.Tokens(ctx, b.DB)
	ok := false

	for _, tok := range toks {
		b.Session.Identify.Token = tok

		err := b.Session.Open()
		if err != nil {
			err = fmt.Errorf("tok %q: %w", tok, err)
			Log.Warn(_f, err)
		} else {
			ok = true
			break
		}
	}

	if !ok {
		return ErrAllToksFailed
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
