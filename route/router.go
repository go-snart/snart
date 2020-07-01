package route

import (
	"strings"

	dg "github.com/bwmarrin/discordgo"
	re2 "github.com/dlclark/regexp2"
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
	_f := "(*Router).Ctx"

	Log.Debugf(_f, "%s %q %q %q", m.ID, line, pfx, cpfx)

	c := &Ctx{}
	c.Prefix = pfx
	c.CleanPrefix = cpfx
	c.Session = s
	c.Message = m

	line = strings.TrimPrefix(line, pfx)
	Log.Debugf(_f, "%s %q %q %q", m.ID, line, pfx, cpfx)
	Log.Debugf(_f, "%#v", *c)

	for _, r := range *rr {
		Log.Debugf(_f, "try route %#v", r)

		if r.match == nil {
			match := `(` + r.Match + `)\b`
			exp, err := re2.Compile(match, re2.IgnoreCase)
			if err != nil {
				Log.Warnf(_f, "re compile %#q: %s", match, err)
				continue
			}
			r.match = exp
		}

		Log.Debugf(_f, "%#v", r.match)

		// can't error - already compiled
		m, _ := r.match.FindStringMatch(line)
		if m == nil {
			Log.Warnf(_f, "re match nil")
			continue
		}
		if m.Index > 0 {
			Log.Warnf(_f, "re match index == %d > 0", m.Index)
			continue
		}

		if r.Okay == nil {
			Log.Warnf(_f, "nil okay, setting to true")
			r.Okay = True
		}

		if r.Okay(c) {
			c.Route = r
			line = line[m.Index:]
			break
		}
	}

	cont := strings.TrimPrefix(line, pfx)
	cont = strings.Trim(cont, " ")
	args := Split(cont)
	Log.Debugf(_f, "%#v", args)

	if len(args) == 0 {
		return nil
	}

	cmd := args[0]
	args = args[1:]

	Log.Debugf(_f, "args %#v", args)

	c.Flags = NewFlags(c, cmd, args)

	Log.Debugf(_f, "ctx %#v", c)

	if c.Route == nil {
		return nil
	}

	return c
}
