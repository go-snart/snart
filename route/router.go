package route

import (
	"fmt"
	re "regexp"
	"strings"

	dg "github.com/bwmarrin/discordgo"
)

type Router []*Route

func NewRouter() *Router {
	rr := make(Router, 0)
	return &rr
}

func (rr *Router) Add(rs ...*Route) {
	*rr = append(*rr, rs...)
}

func (rr *Router) Ctx(pfx, cpfx string, s *dg.Session, m *dg.Message, line string) *Ctx {
	_f := "(Router).Ctx"

	Log.Debugf(_f, "%s %s %s %s", m.GuildID, line, pfx, cpfx)

	c := &Ctx{}
	c.Prefix = pfx
	c.CleanPrefix = cpfx
	c.Session = s
	c.Message = m

	cont := strings.TrimPrefix(line, pfx)
	cont = strings.Trim(cont, " ")
	args := Split(cont)
	Log.Debugf(_f, "%#v", args)

	if len(args) == 0 {
		return nil
	}

	cmd := args[0]
	args = args[1:]

	Log.Debug(_f, "args", args)

	c.Flags = NewFlags(c, cmd, args)

	for _, r := range *rr {
		Log.Debug(_f, "try route", r)

		Log.Debugf(_f, "flag name %s", c.Flags.Name())

		matched, err := re.MatchString("^"+r.Match+"$", c.Flags.Name())
		if err != nil {
			err = fmt.Errorf("re match %#v %#v: %w", r.Match, c.Flags.Name(), err)
			Log.Warn(_f, err)
			return nil
		}

		var ok bool
		if r.Okay == nil {
			ok = true
		} else {
			ok = r.Okay(c)
		}

		if matched && ok {
			c.Route = r
			break
		}
	}

	Log.Debugf(_f, "ctx %#v", c)

	if c.Route == nil {
		return nil
	}

	return c
}
