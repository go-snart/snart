// Package admin contains a plugin that provides admin-only commands for Snart Bots.
package admin

import "github.com/go-snart/snart/bot"

// Admin is a plugin that provides admin-only commands for Snart Bots.
type Admin struct {
	Bot *bot.Bot
}

// String returns the Admin's string representation.
func (a *Admin) String() string {
	return "builtin admin plug"
}
