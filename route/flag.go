package route

import (
	"flag"
	"fmt"

	dg "github.com/bwmarrin/discordgo"

	"github.com/go-snart/snart/log"
)

// Flag holds a FlagSet for a Ctx.
type Flag struct {
	*flag.FlagSet
	args []string
	ctx  *Ctx
	err  error
}

// NewFlag creates a Flag.
func NewFlag(ctx *Ctx, name string, args []string) *Flag {
	f := &Flag{
		FlagSet: flag.NewFlagSet(name, flag.ContinueOnError),
		ctx:     ctx,
		args:    args,
	}

	f.FlagSet.Usage = func() {
		err := f.Usage().Send()
		if err != nil {
			log.Warn.Println(err)
		}
	}

	return f
}

// Usage generates a *Reply containing usage info from the Flag.
func (f *Flag) Usage() *Reply {
	rep := f.ctx.Reply()
	if f.err != nil {
		rep.Content = "**Error:** " + f.err.Error()
	}

	rep.Embed = &dg.MessageEmbed{
		Title:       "Usage of `" + f.ctx.Route.Name + "`",
		Description: f.ctx.Route.Desc,
		Fields:      make([]*dg.MessageEmbedField, 0),
	}

	f.VisitAll(func(f *flag.Flag) {
		field := &dg.MessageEmbedField{
			Name:  fmt.Sprintf("Flag `-%s`", f.Name),
			Value: fmt.Sprintf("(default `%s`)\n%s", f.DefValue, f.Usage),
		}
		rep.Embed.Fields = append(rep.Embed.Fields, field)
	})

	return rep
}

// Parse parses the arguments given to the Flag.
func (f *Flag) Parse() error {
	return f.FlagSet.Parse(f.args)
}
