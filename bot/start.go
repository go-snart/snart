package bot

import (
	"fmt"
	"time"
)

// Start performs the Bot's startup functions, and then waits until an interrupt.
func (b *Bot) Start() error {
	_f := "(*Bot).Start"

	b.GoPlugins()

	b.Startup = time.Now()
	Log.Infof(_f, "startup at %s", b.Startup)

	err := b.DB.Start()
	if err != nil {
		err = fmt.Errorf("db start: %w", err)
		Log.Error(_f, err)

		return err
	}

	Log.Info(_f, "db started")

	tok, err := b.DB.Token()
	if err != nil {
		err = fmt.Errorf("token: %w", err)
		Log.Error(_f, err)

		return err
	}

	Log.Info(_f, "got token")

	b.Session.Token = tok.Value

	err = b.Session.Open()
	if err != nil {
		err = fmt.Errorf("session open: %w", err)
		Log.Error(_f, err)

		return err
	}

	Log.Info(_f, "session opened")

	go b.CycleGamers()

	b.WaitReady()
	Log.Info(_f, "ready")

	b.HandleInterrupts()

	b.Logout()

	return nil
}

func (b *Bot) Logout() {
	_f := "(*Bot).Logout"

	err := b.Session.Close()
	if err != nil {
		err = fmt.Errorf("close: %w", err)
		Log.Warn(_f, err)
	}

	Log.Info(_f, "logged out")
}
