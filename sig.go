package bot

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type SigError struct {
	os.Signal
}

func (s SigError) Error() string {
	return fmt.Sprintf("sig: %s", s.Signal)
}

// SigChan prepares an error channel that listens for os.Signals and context cancellation.
func SigChan(ctx context.Context) chan error {
	ch := make(chan error, 1)
	sig := make(chan os.Signal, 1)

	signal.Notify(sig, os.Interrupt)
	signal.Notify(sig, syscall.SIGTERM)

	go func() {
		select {
		case s := <-sig:
			ch <- SigError{s}
		case <-ctx.Done():
			ch <- ctx.Err()
		}
	}()

	return ch
}
