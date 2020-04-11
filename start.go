package bot

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-snart/snart/lib/errs"
)

func (b *Bot) Start() error {
	_f := "(*Bot).Start"

	b.Startup = time.Now()
	Log.Infof(_f, "startup at %s", b.Startup)

	err := b.DB.Start()
	if err != nil {
		errs.Wrap(&err, `b.DB.Start()`)
		Log.Error(_f, err)
		return err
	}
	Log.Info(_f, "db started")

	tok, err := b.Token()
	if err != nil {
		errs.Wrap(&err, `b.Token()`)
		Log.Error(_f, err)
		return err
	}
	Log.Info(_f, "get token")
	b.Session.Token = tok.Value

	err = b.Session.Open()
	if err != nil {
		errs.Wrap(&err, `b.Session.Open()`)
		Log.Error(_f, err)
		return err
	}
	Log.Info(_f, "session opened")

	b.WaitReady()
	Log.Info(_f, "ready")

	signal.Notify(b.Sig, os.Interrupt)
	signal.Notify(b.Sig, syscall.SIGTERM)
	<-b.Sig
	Log.Info(_f, "exiting")

	return nil
}

func (b *Bot) Uptime() time.Duration {
	return time.Since(b.Startup)
}
