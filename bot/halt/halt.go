package halt

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// Halt wraps an os.Signal or error which causes a Bot to exit.
type Halt struct {
	Sig os.Signal
	Err error
}

// Error returns a string version of the Halt's error message, or "unknown" if it's indeterminable.
func (h Halt) Error() string {
	switch {
	case h.Sig != nil:
		return fmt.Sprintf("sig: %s", h.Sig)
	case h.Err != nil:
		return fmt.Sprintf("err: %s", h.Err)
	default:
		return "unknown"
	}
}

// Unwrap returns the underlying error, if there is one.
func (h Halt) Unwrap() error {
	if h.Err != nil {
		return h.Err
	}

	return nil
}

// Chan prepares a Halt channel with running signal notifications.
func Chan(ctx context.Context) chan Halt {
	halts := make(chan Halt, 1)
	sig := make(chan os.Signal, 1)

	signal.Notify(sig, os.Interrupt)
	signal.Notify(sig, syscall.SIGTERM)

	go func() {
		for {
			select {
			case s := <-sig:
				halts <- Halt{Sig: s}
			case <-ctx.Done():
				halts <- Halt{Err: ctx.Err()}
			}
		}
	}()

	return halts
}
