package route

import (
	"errors"
	"flag"
	"fmt"
	"strings"

	dg "github.com/bwmarrin/discordgo"
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
	f := &Flag{}

	f.ctx = ctx
	f.args = args

	f.FlagSet = flag.NewFlagSet(name, flag.ContinueOnError)
	f.FlagSet.Usage = func() {
		err := f.Usage().Send()
		if err != nil {
			Warn.Println(err)
		}
	}
	f.FlagSet.SetOutput(&strings.Builder{})

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
			Name:  "Flag `-" + f.Name + "`",
			Value: f.Usage,
		}
		rep.Embed.Fields = append(rep.Embed.Fields, field)
	})

	return rep
}

// Parse parses the arguments given to the Flag.
func (f *Flag) Parse() error {
	err := f.FlagSet.Parse(f.args)
	if err != nil {
		f.err = err

		if errors.Is(err, flag.ErrHelp) {
			return err
		}

		err = fmt.Errorf("flag parse %#v: %w", f.args, err)
		Warn.Println(err)

		return err
	}

	return nil
}

// Output retrieves the Flag's FlagSet's Output as a string.
func (f *Flag) Output() string {
	if b, ok := f.FlagSet.Output().(fmt.Stringer); ok {
		return b.String()
	}

	return ""
}
