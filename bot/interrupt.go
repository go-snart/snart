package bot

import (
	"os"
	"os/signal"
	"syscall"
)

// Interrupt wraps an os.Signal or error which causes a Bot to exit.
type Interrupt struct {
	Sig os.Signal
	Err error
}

// NotifyInterrupt uses signal.Notify and a pipe goroutine to send OS-level interrupts on the Bot's Interrupt.
func (b *Bot) NotifyInterrupt(sigs ...os.Signal) {
	sig := make(chan os.Signal)

	go func(sig chan os.Signal, interrupt chan Interrupt) {
		for s := range sig {
			interrupt <- Interrupt{Sig: s}
		}
	}(sig, b.Interrupt)

	signal.Notify(sig, os.Interrupt)
	signal.Notify(sig, syscall.SIGTERM)
}
