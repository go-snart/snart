// Package bot contains the general workings of a Snart Bot.
package bot

import (
	"fmt"
	"log"
	"time"

	"github.com/diamondburned/arikawa/v2/gateway"
	"github.com/diamondburned/arikawa/v2/state"

	"github.com/go-snart/db"
	"github.com/go-snart/route"
)

// DefaultIntents are the basic Intents a Bot needs.
const DefaultIntents = gateway.IntentGuilds

/*
const DefaultIntents = gateway.IntentGuilds |
	gateway.IntentGuildBans |
	gateway.IntentGuildEmojis |
	gateway.IntentGuildIntegrations |
	gateway.IntentGuildWebhooks |
	gateway.IntentGuildInvites |
	gateway.IntentGuildVoiceStates |
	gateway.IntentGuildMessages |
	gateway.IntentGuildMessageReactions |
	gateway.IntentGuildMessageTyping |
	gateway.IntentDirectMessages |
	gateway.IntentDirectMessageReactions |
	gateway.IntentDirectMessageTyping
*/

// Bot holds all the workings of a Snart bot.
type Bot struct {
	DB      *db.DB
	State   *state.State
	Route   *route.Route
	Intents gateway.Intents
	Gamers  []Gamer
	Errs    chan error
}

// Open creates a Bot from the given database uri.
func Open(uri string) (*Bot, error) {
	log.Printf("opening %s", uri)

	d, err := db.Open(uri)
	if err != nil {
		return nil, fmt.Errorf("db open %q: %w", uri, err)
	}
	log.Printf("opened db %s", uri)

	return New(d)
}

// New creates a Bot from the given DB.
func New(d *db.DB) (*Bot, error) {
	log.Printf("new from db %s", d)

	tok, err := getToken(d)
	if err != nil {
		return nil, fmt.Errorf("get tok: %w", err)
	}
	log.Printf("got token %q", tok)

	s, err := state.New(tok)
	if err != nil {
		// don't log the token
		return nil, fmt.Errorf("new state: %w", err)
	}
	log.Printf("made state")

	r := route.New(d, s)
	log.Printf("made route")

	return &Bot{
		DB:      d,
		State:   s,
		Route:   r,
		Intents: DefaultIntents,
		Gamers: []Gamer{
			GamerTimer(time.Now()), // uptime
		},
		Errs: make(chan error),
	}, nil
}

func getToken(d *db.DB) (string, error) {
	log.Printf("getting token")

	const key = "token"

	tok := ""

	err := d.Get(key, &tok)
	if err != nil {
		return "", fmt.Errorf("db get %q: %w", key, err)
	}
	log.Printf("got token %q from db", tok)

	return tok, nil
}
