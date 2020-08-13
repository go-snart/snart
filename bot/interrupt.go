package bot

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-snart/snart/logs"
)

// Interrupt wraps an os.Signal or error which causes a Bot to exit.
type Interrupt struct {
	Sig os.Signal
	Err error
}

// Error returns a string version of the Interrupt's error message, or "unknown" if it's indeterminable.
func (i Interrupt) Error() string {
	switch {
	case i.Sig != nil:
		return fmt.Sprintf("sig: %s", i.Sig)
	case i.Err != nil:
		return fmt.Sprintf("err: %s", i.Err)
	default:
		return "unknown"
	}
}

// Unwrap returns the underlying error, if there is one.
func (i Interrupt) Unwrap() error {
	if i.Err != nil {
		return i.Err
	}

	return nil
}

func (b *Bot) handleInterrupts() {
	sig := make(chan os.Signal, 1)

	go func(sig chan os.Signal, interrupt chan Interrupt) {
		for s := range sig {
			fmt.Println()
			interrupt <- Interrupt{Sig: s}
		}
	}(sig, b.Interrupt)

	signal.Notify(sig, os.Interrupt)
	signal.Notify(sig, syscall.SIGTERM)

	err := fmt.Errorf("interrupt: %w", <-b.Interrupt)
	logs.Warn.Println(err)
}
