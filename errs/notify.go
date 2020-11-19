package errs

import (
	"os"
	"os/signal"
	"syscall"
)

// Notify prepares an error channel to listen for Signals.
func Notify(ch chan error) {
	ss := make(chan os.Signal, 1)

	signal.Notify(ss, os.Interrupt)
	signal.Notify(ss, syscall.SIGTERM)

	go func() {
		ch <- Error{
			Sig: <-ss,
		}
	}()
}
