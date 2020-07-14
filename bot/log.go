package bot

import (
	"fmt"

	dg "github.com/bwmarrin/discordgo"
	"github.com/superloach/minori"
)

// Log is the logger for the bot package.
var Log = minori.GetLogger("bot")

var _ = func() interface{} {
	dg.Logger = DiscordGoLogger
	return nil
}()

// DiscordGoLogger replaces dg's Logger with minori.
func DiscordGoLogger(msgL, caller int, format string, a ...interface{}) {
	var lvl rune

	switch msgL {
	case dg.LogError:
		lvl = 'e'
	case dg.LogWarning:
		lvl = 'w'
	case dg.LogInformational:
		lvl = 'i'
	case dg.LogDebug:
		lvl = 'd'
	default:
		lvl = '?'
	}

	_f := fmt.Sprintf(
		"[dg:%c:%d]",
		lvl, caller,
	)

	Log.Debugf(_f, format, a...)
}
