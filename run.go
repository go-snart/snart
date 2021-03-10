package snart

import (
	"fmt"
	"log"
	"os"
	"os/signal"
)

// sigError occurs when a process signal kills the bot.
type sigError struct {
	Signal os.Signal
}

// Error implements error.
func (s sigError) Error() string {
	return "killed by signal " + s.Signal.String()
}

// Run performs the Bot's startup functions, and waits for an error.
func (b *Bot) Run() error {
	log.Println("running bot")

	sigs := make(chan os.Signal, 1)

	go func() {
		b.Errs <- sigError{
			Signal: <-sigs,
		}
	}()

	signal.Notify(sigs, os.Interrupt)
	log.Println("listening for sigs")

	b.State.AddHandler(b.Route.Handle)
	log.Println("router handler added")

	b.State.Gateway.Identifier.Intents = b.Intents

	log.Println("intents injected")

	err := b.State.Open()
	if err != nil {
		return fmt.Errorf("open state: %w", err)
	}

	log.Println("state opened")

	defer b.State.Close()

	go b.CycleGamers()
	log.Println("gamers cycling")

	err = <-b.Errs
	if err != nil {
		return fmt.Errorf("from errs chan: %w", err)
	}

	return nil
}
