package help

import (
	"fmt"
	"sort"
	"strings"

	dg "github.com/bwmarrin/discordgo"

	"github.com/go-snart/snart/log"
	"github.com/go-snart/snart/route"
)

func (h *Help) routesByCat(c *route.Ctx) map[string][]*route.Route {
	byCat := map[string][]*route.Route{}

	for _, r := range h.Handler.Routes {
		if r.Okay == nil {
			r.Okay = route.True
		}

		if r.Okay(c) {
			cat := r.Cat
			if cat == "" {
				cat = "misc"
			}

			byCat[cat] = append(byCat[cat], r)
		}
	}

	return byCat
}

// Menu gives a help menu.
func (h *Help) Menu(c *route.Ctx) error {
	err := c.Flag.Parse()
	if err != nil {
		err = fmt.Errorf("flag parse: %w", err)
		log.Warn.Println(err)

		return err
	}

	args := c.Flag.Args()
	if len(args) > 0 {
		cmd := args[0]

		hc := h.Handler.Ctx(c, c.Prefix, c.Session, nil, c.Prefix.Value+cmd+" -help")
		if hc == nil {
			rep := c.Reply()
			rep.Content = fmt.Sprintf("command `%s` not known", cmd)

			return rep.Send()
		}

		err = hc.Run()
		if err != nil {
			err = fmt.Errorf("hc run: %w", err)
			log.Warn.Println(err)

			return err
		}
	}

	pfx := c.Prefix.Clean

	rep := c.Reply()
	rep.Embed = &dg.MessageEmbed{
		Title:       c.Session.State.User.Username + " Help",
		Description: fmt.Sprintf("Prefix: `%s`", pfx),
	}

	byCat := h.routesByCat(c)

	cats := make([]string, 0)

	for cat, rs := range byCat {
		if len(rs) > 0 {
			cats = append(cats, cat)
		}
	}

	sort.Strings(cats)

	for _, name := range cats {
		rs, ok := byCat[name]
		if !ok {
			continue
		}

		addCatField(rep.Embed, name, rs, pfx)
	}

	rep.Embed.Footer = &dg.MessageEmbedFooter{
		Text: "use the `-help` flag on a command for detailed help",
	}

	return rep.Send()
}

func addCatField(e *dg.MessageEmbed, name string, rs []*route.Route, pfx string) {
	field := &dg.MessageEmbedField{
		Name:   name,
		Inline: false,
	}

	for _, r := range rs {
		desc := r.Desc
		lines := strings.Split(desc, "\n")

		if len(lines) == 0 {
			desc = "*no description*"
		} else {
			desc = lines[0]
		}

		field.Value += fmt.Sprintf("`%s%s`: %s\n", pfx, r.Name, desc)
	}

	e.Fields = append(e.Fields, field)
}
