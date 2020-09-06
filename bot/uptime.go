package bot

import (
	"fmt"
	"time"

	dg "github.com/bwmarrin/discordgo"
)

// Uptime is a gamer.Gamer that shows the Bot's uptime.
func (b *Bot) Uptime() *dg.Game {
	return &dg.Game{
		Name: fmt.Sprintf("for %s", time.Since(b.Startup).Round(time.Second)),
		Type: dg.GameTypeGame, // playing...
	}
}
