package help

import (
	"fmt"
	"sort"
	"strings"

	dg "github.com/bwmarrin/discordgo"

	"github.com/go-snart/snart/log"
	"github.com/go-snart/snart/route"
)

func (h *Help) routesByCat(ctx *route.Ctx) map[string][]*route.Route {
	byCat := map[string][]*route.Route{}

	for _, r := range *h.Handler {
		if r.Okay == nil {
			r.Okay = route.True
		}

		if r.Okay(ctx) {
			cat := r.Cat
			if cat == "" {
				cat = "misc"
			}

			byCat[cat] = append(byCat[cat], r)
		}
	}

	return byCat
}

// Help gives a help menu.
func (h *Help) Menu(ctx *route.Ctx) error {
	err := ctx.Flag.Parse()
	if err != nil {
		err = fmt.Errorf("flag parse: %w", err)
		log.Warn.Println(err)

		return err
	}

	args := ctx.Flag.Args()
	if len(args) > 0 {
		cmd := args[0]

		rctx := h.Handler.Ctx(ctx.Prefix, ctx.Session, nil, ctx.Prefix.Value+cmd+" -help")
		if rctx == nil {
			rep := ctx.Reply()
			rep.Content = fmt.Sprintf("command `%s` not known", cmd)
			return rep.Send()
		}

		err = rctx.Run()
		if err != nil {
			err = fmt.Errorf("rctx run: %w", err)
			log.Warn.Println(err)
			return err
		}
	}

	pfx := ctx.Prefix.Clean

	rep := ctx.Reply()
	rep.Embed = &dg.MessageEmbed{
		Title:       ctx.Session.State.User.Username + " Help",
		Description: fmt.Sprintf("Prefix: `%s`", pfx),
	}

	byCat := h.routesByCat(ctx)

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
