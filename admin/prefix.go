package admin

import (
	"errors"
	"flag"
	"fmt"

	"github.com/go-snart/snart/log"
	"github.com/go-snart/snart/route"
)

// Prefix is a command that can change prefixes.
func (a *Admin) Prefix(c *route.Ctx) error {
	guild := c.Flag.String("guild", c.Message.GuildID, "guild id to view/change prefix for")
	// set := c.Flag.String("set", "", "new prefix value to set")

	err := c.Flag.Parse()
	if err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return nil
		}

		err = fmt.Errorf("flag parse: %w", err)
		log.Warn.Println(err)

		return err
	}

	gid := *guild

	if gid == "default" {
		gid = ""
	}

	pfx, err := a.DB.GuildPrefix(c, gid)
	if err != nil {
		err = fmt.Errorf("guild prefix: %w", err)
		log.Warn.Println(err)

		return err
	}

	rep := c.Reply()

	if pfx == nil {
		rep.Content = fmt.Sprintf("no prefix for guild `%s`", *guild)
	} else {
		rep.Content = fmt.Sprintf("prefix for guild `%s` is `%s`", *guild, pfx.Value)
	}

	return rep.Send()
}
