package route

import (
	"flag"
	"fmt"
	"strings"

	dg "github.com/bwmarrin/discordgo"
)

type Flags struct {
	*flag.FlagSet
	args []string
	ctx  *Ctx
	err  error
}

func NewFlags(ctx *Ctx, name string, args []string) *Flags {
	f := &Flags{}

	f.ctx = ctx
	f.args = args

	f.FlagSet = flag.NewFlagSet(name, flag.ContinueOnError)
	f.FlagSet.Usage = func() {
		f.Usage().Send()
	}
	f.FlagSet.SetOutput(&strings.Builder{})

	return f
}

func visitor(rep *Reply) func(f *flag.Flag) {
	return func(f *flag.Flag) {
		field := &dg.MessageEmbedField{
			Name:  "Flag `-" + f.Name + "`",
			Value: f.Usage,
		}
		rep.Embed.Fields = append(rep.Embed.Fields, field)
	}
}

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
	f.VisitAll(visitor(rep))

	return rep
}

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

func (f *Flags) Output() string {
	if b, ok := f.FlagSet.Output().(interface {
		String() string
	}); ok {
		return b.String()
	}
	return ""
}
