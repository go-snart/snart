package admin

import (
	"errors"
	"flag"
	"fmt"

	"github.com/go-snart/snart/log"
	"github.com/go-snart/snart/route"
)

// Prefix is a command that can change prefixes.
func (a *Admin) Prefix(ctx *route.Ctx) error {
	guild := ctx.Flag.String("guild", ctx.Message.GuildID, "guild to change the prefix for")
	value := ctx.Flag.String("value", ctx.Prefix.Value, "prefix value to use")

	err := ctx.Flag.Parse()
	if err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return nil
		}

		err = fmt.Errorf("flag parse: %w", err)
		log.Warn.Println(err)

		return err
	}

	rep := ctx.Reply()
	rep.Content = fmt.Sprintf("guild=%#q value=%#q", *guild, *value)

	return rep.Send()
}
