package bot

import (
	"fmt"
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
	sig := make(chan os.Signal, 1)

	go func(sig chan os.Signal, interrupt chan Interrupt) {
		for s := range sig {
			interrupt <- Interrupt{Sig: s}
		}
	}(sig, b.Interrupt)

	signal.Notify(sig, os.Interrupt)
	signal.Notify(sig, syscall.SIGTERM)
}

// HandleInterrupts sets up and handles Interrupts for the Bot.
func (b *Bot) HandleInterrupts() {
	b.NotifyInterrupt()

	interrupt := <-b.Interrupt
	err := "interrupt: unknown"

	switch {
	case interrupt.Err != nil:
		err = fmt.Sprintf("interrupt: err %s", interrupt.Err)
	case interrupt.Sig != nil:
		err = fmt.Sprintf("interrupt: sig %s", interrupt.Sig)
	}

	warn.Println(err)
}
