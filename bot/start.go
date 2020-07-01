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
	Log.Info(_f, "get token")
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

	b.NotifyInterrupt()
	interrupt := <-b.Interrupt
	if interrupt.Err != nil {
		err = fmt.Errorf("interrupt: %w", interrupt.Err)
	} else if interrupt.Sig != nil {
		err = fmt.Errorf("interrupt: %s", interrupt.Sig)
	} else {
		err = fmt.Errorf("interrupt: unknown")
	}
	Log.Error(_f, err)

	if !b.Session.State.User.Bot {
		err = b.Session.Logout()
		if err != nil {
			err = fmt.Errorf("logout: %w", err)
			Log.Error(_f, err)
			return err
		}

		Log.Info(_f, "logged out")
		return nil
	}

	err = b.Session.Close()
	if err != nil {
		err = fmt.Errorf("close: %w", err)
		Log.Error(_f, err)
		return err
	}

	Log.Info(_f, "session closed")
	return nil
}
