package route

import (
	"flag"
	"fmt"
	"strings"

	dg "github.com/bwmarrin/discordgo"
)

// Flags holds a FlagSet for a Ctx.
type Flags struct {
	*flag.FlagSet
	args []string
	ctx  *Ctx
	err  error
}

// NewFlags creates a Flags.
func NewFlags(ctx *Ctx, name string, args []string) *Flags {
	f := &Flags{}

	f.ctx = ctx
	f.args = args

	f.FlagSet = flag.NewFlagSet(name, flag.ContinueOnError)
	f.FlagSet.Usage = func() {
		_f := "f.Flagset.Usage"

		err := f.Usage().Send()
		if err != nil {
			Log.Warn(_f, err)
		}
	}
	f.FlagSet.SetOutput(&strings.Builder{})

	return f
}

// Usage generates a *Reply containing usage info from the Flags.
func (f *Flags) Usage() *Reply {
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

// Parse parses the arguments given to the Flags.
func (f *Flags) Parse() error {
	_f := "(*Flags).Parse"

	err := f.FlagSet.Parse(f.args)
	if err != nil {
		f.err = err

		err = fmt.Errorf("flag parse %#v: %w", f.args, err)
		Log.Error(_f, err)

		return err
	}

	return nil
}

// Output retrieves the Flags' FlagSet's Output as a string.
func (f *Flags) Output() string {
	if b, ok := f.FlagSet.Output().(fmt.Stringer); ok {
		return b.String()
	}

	return ""
}
