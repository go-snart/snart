package admin

import (
	"log"

	"github.com/go-snart/route"
	"github.com/go-snart/snart"
)

// Plug is the default instance of Admin.
//nolint:gochecknoglobals // plug convention
var Plug = &Admin{
	Errs: nil,
}

// Plug connects the Admin with the given Bot.
func (a *Admin) Plug(b *snart.Bot) error {
	log.Println("plugging admin")

	a.Errs = b.Errs

	log.Println("set bot")

	b.Route.AddCmds(route.Cmd{
		Name:  "restart",
		Desc:  "restart the bot",
		Cat:   a.String(),
		Func:  a.Restart,
		Hide:  true,
		Flags: nil,
	})

	log.Println("added routes")

	return nil
}
