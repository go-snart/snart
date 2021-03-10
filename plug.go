package snart

import (
	"fmt"
	"log"
)

// Plug describes a plugin for a Bot.
type Plug interface {
	fmt.Stringer

	Plug(*Bot) error
}

// Plug applies the given Plug to the Bot.
func (b *Bot) Plug(p Plug) error {
	log.Printf("plugging %s", p)

	return p.Plug(b)
}
