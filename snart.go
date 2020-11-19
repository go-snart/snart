// Package bot contains the general workings of a Snart Bot.
package snart

import (
	"fmt"
	"time"

	"github.com/diamondburned/arikawa/gateway"
	"github.com/diamondburned/arikawa/state"

	"github.com/go-snart/db"
	"github.com/go-snart/logs"
	"github.com/go-snart/plug/gamer"
	"github.com/go-snart/route"
)

// DefaultIntents are the unprivileged intents used by default.
//
// Omitted intents include:
//  gateway.IntentGuildMembers
//  gateway.IntentGuildPresences
var DefaultIntents = gateway.IntentGuilds | gateway.IntentGuildBans | gateway.IntentGuildEmojis |
	gateway.IntentGuildIntegrations | gateway.IntentGuildWebhooks | gateway.IntentGuildInvites |
	gateway.IntentGuildVoiceStates | gateway.IntentGuildMessages | gateway.IntentGuildMessageReactions |
	gateway.IntentGuildMessageTyping | gateway.IntentDirectMessages | gateway.IntentDirectMessageReactions |
	gateway.IntentDirectMessageTyping

// Snart holds all the workings of a Snart bot.
type Snart struct {
	DB    *db.DB
	State *state.State
	Route *route.Route
	Err   chan error

	Intents gateway.Intents
	Gamers  gamer.Queue
}

// Open creates a Snart.
func Open(uri string) (*Snart, error) {
	d, err := db.Open(uri)
	if err != nil {
		return nil, fmt.Errorf("db open %q: %w", err)
	}

	return New(d)
}

// New creates a Snart from the given DB.
func New(d *db.DB) (*Snart, error) {
	s := &Snart{
		DB:      d,
		State:   nil,
		Route:   nil,
		Err:     make(chan error),
		Intents: DefaultIntents,
		Gamers: gamer.Queue{
			gamer.Uptime(time.Now()),
		},
	}

	tok, err := s.Token()
	if err != nil {
		return nil, err
	}

	state, err := state.NewWithIntents(tok, b.Intents)
	if err != nil {
		return nil, fmt.Errorf("new state %q: %w", tok, err)
	}

	ready := make(chan struct{})

	rm := state.AddHandler(func(_ *gateway.ReadyEvent) {
		logs.Info.Println("ready")
		close(ready)
	})

	err = state.Open()
	if err != nil {
		return nil, fmt.Errorf("open %q: %w", tok, err)
	}

	<-ready
	rm()

	return s, nil
}

func (s *Snart) Token() (string, error) {
	const key = "token"

	tok := ""

	err := s.DB.Get(key, &tok)
	if err != nil {
		return "", fmt.Errorf("db get %q: %w", key, err)
	}

	return tok, nil
}
