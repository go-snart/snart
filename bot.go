// Package snart contains the general workings of a Snart Bot.
package snart

import (
	"fmt"
	"log"
	"time"

	"github.com/diamondburned/arikawa/v2/gateway"
	"github.com/diamondburned/arikawa/v2/state"
	"github.com/superloach/confy"

	"github.com/go-snart/route"
)

// KeyToken is the Confy key used to fetch the Bot's token.
const KeyToken = "token"

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
func New(c confy.Confy) (*Bot, error) {
	token := ""

	err := c.Get(KeyToken, &token)
	if err != nil {
		return nil, fmt.Errorf("load %q: %w", KeyToken, err)
	}

	s, err := state.New(token)
	if err != nil {
		return nil, fmt.Errorf("new state %q: %w", token, err)
	}

	log.Println("made state")

	r, err := route.New(s, c)
	if err != nil {
		return nil, fmt.Errorf("new route: %w", err)
	}

	return &Bot{
		Route: r,

		Intents: BaseIntents,
		Gamers: []Gamer{
			GamerTimer(time.Now()), // uptime
		},
		Errs: make(chan error, 1),
	}, nil
}
