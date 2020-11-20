// Package bot contains the general workings of a Snart Bot.
package snart

import (
	"fmt"
	"time"

	"github.com/diamondburned/arikawa/gateway"
	"github.com/diamondburned/arikawa/state"

	"github.com/go-snart/db"
	"github.com/go-snart/route"
	"github.com/go-snart/snart/gamer"
	"github.com/go-snart/snart/logs"
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

// Bot holds all the workings of a Snart bot.
type Bot struct {
	DB    *db.DB
	State *state.State
	Route *route.Route
	Err   chan error

	Intents gateway.Intents
	Gamers  gamer.Queue
}

// Open creates a Bot from the given database uri.
func Open(uri string) (*Bot, error) {
	d, err := db.Open(uri)
	if err != nil {
		return nil, fmt.Errorf("db open %q: %w", err)
	}

	return New(d)
}

// New creates a Bot from the given DB.
func New(d *db.DB) (*Bot, error) {
	b := &Bot{
		DB:      d,
		State:   nil,
		Route:   nil,
		Err:     make(chan error),
		Intents: DefaultIntents,
		Gamers: gamer.Queue{
			gamer.Uptime(time.Now()),
		},
	}

	tok, err := b.Token()
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

	return b, nil
}

func (b *Bot) Token() (string, error) {
	const key = "token"

	tok := ""

	err := b.DB.Get(key, &tok)
	if err != nil {
		return "", fmt.Errorf("db get %q: %w", key, err)
	}

	return tok, nil
}
