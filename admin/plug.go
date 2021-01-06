package admin

import (
	"log"

	"github.com/go-snart/route"
	"github.com/go-snart/snart/bot"
)

// Plug is the default instance of Admin.
//nolint:gochecknoglobals // plug convention
var Plug = &Admin{
	Bot: nil,
}

// Plug injects the Admin into the given Bot.
func (a *Admin) Plug(b *bot.Bot) error {
	log.Println("injecting admin")

	a.Bot = b
	log.Println("set bot")

	b.Route.Add("admin", &route.Command{
		Name:  "restart",
		Desc:  "restart the bot",
		Func:  a.Restart,
		Hide:  true,
		Flags: nil,
	})
	log.Println("added routes")

	return nil
}
