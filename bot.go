// Package snart contains the general workings of a Snart Bot.
package snart

import (
	"fmt"
	"log"
	"time"

	"github.com/diamondburned/arikawa/v2/gateway"
	"github.com/diamondburned/arikawa/v2/state"

	"github.com/go-snart/route"
)

// BaseIntents is the basic intents needed by a Bot.
const BaseIntents = gateway.IntentGuildMessages

// Bot holds all the workings of a Snart bot.
type Bot struct {
	*route.Route

	Intents gateway.Intents
	Gamers  []Gamer
	Errs    chan error
}

// New creates a Bot with the given token.
func New(token string) (*Bot, error) {
	s, err := state.New(token)
	if err != nil {
		return nil, fmt.Errorf("new state %q: %w", token, err)
	}

	log.Println("made state")

	return &Bot{
		Route: route.New(s),

		Intents: BaseIntents,
		Gamers: []Gamer{
			GamerTimer(time.Now()), // uptime
		},
		Errs: make(chan error, 1),
	}, nil
}
