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

func MkFlags(ctx *Ctx, name string, args []string) *Flags {
	f := &Flags{}

	f.ctx = ctx
	f.args = args

	f.FlagSet = flag.NewFlagSet(name, flag.ContinueOnError)
	f.FlagSet.Usage = f.Usage
	f.FlagSet.SetOutput(&strings.Builder{})

	return f
}

func (f *Flags) Usage() {
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

	rep.Send()
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

type DgUserValue struct {
	user *dg.User
	ctx  *Ctx
}

func (d *DgUserValue) String() string {
	return d.user.String()
}

func (d *DgUserValue) Set(s string) error {
	g, err := d.ctx.Session.Guild(d.ctx.Message.GuildID)
	if err != nil {
		return err
	}

	for _, m := range g.Members {
		if m.Mention() == s ||
			m.Nick == s ||
			m.User.Mention() == s ||
			m.User.String() == s ||
			m.User.Username == s ||
			m.User.Username == s ||
			m.User.ID == s {
			d.user = m.User
			return nil
		}
	}

	return nil
}

func (f *Flags) User(name string, value *dg.User, usage string) *dg.User {
	val := &DgUserValue{
		user: value,
		ctx:  f.ctx,
	}

	f.Var(val, name, usage)

	return val.user
}
