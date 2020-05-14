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

func (rr *Router) Ctx(pfx, cpfx string, s *dg.Session, m *dg.Message, line string) (*Ctx, error) {
	_f := "(Router).Ctx"

	Log.Debugf(_f, "%s %s %s %s", m.GuildID, line, pfx, cpfx)

	c := &Ctx{}
	c.Prefix = pfx
	c.CleanPrefix = cpfx
	c.Session = s
	c.Message = m

	cont := strings.TrimPrefix(line, pfx)
	cont = strings.Trim(cont, " ")
	args, err := Split(cont)
	if err != nil {
		err = fmt.Errorf("split %#v: %w", cont, err)
		Log.Error(_f, err)
		return nil, err
	}
	Log.Debugf(_f, "%#v", args)

	if len(args) == 0 {
		return nil, nil
	}

	cmd := args[0]

	if len(args) == 1 {
		args = make([]string, 0)
	} else {
		args = args[1:]
	}

	Log.Debug(_f, "args", args)

	c.Flags = MkFlags(c, cmd, args)

	for _, r := range *rr {
		Log.Debug(_f, "try route", r)

		Log.Debugf(_f, "flag name %s", c.Flags.Name())

		matched, err := re.MatchString("^"+r.Match+"$", c.Flags.Name())
		if err != nil {
			err = fmt.Errorf("re match %#v %#v: %w", r.Match, c.Flags.Name(), err)
			Log.Error(_f, err)
			return nil, err
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
		return nil, nil
	}

	return c, nil
}
