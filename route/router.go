package route

import (
	"strings"

	dg "github.com/bwmarrin/discordgo"
	re2 "github.com/dlclark/regexp2"
)

// Router is a slice of Routes.
type Router []*Route

// NewRouter creates a Router.
func NewRouter() *Router {
	rr := make(Router, 0)
	return &rr
}

// Add adds a Route to the Router.
func (rr *Router) Add(rs ...*Route) {
	*rr = append(*rr, rs...)
}

// Ctx gets a Ctx by finding an appropriate Route for a given prefix, session, message, etc.
func (rr *Router) Ctx(pfx, cpfx string, s *dg.Session, m *dg.Message, line string) *Ctx {
	_f := "(*Router).Ctx"

	c := &Ctx{Prefix: pfx, CleanPrefix: cpfx, Session: s, Message: m}

	line = strings.TrimSpace(strings.TrimPrefix(line, pfx))

	for _, r := range *rr {
		Log.Debugf(_f, "try route %#v", r)

		if r.match == nil {
			match := `(` + r.Match + `)\b`

			exp, err := re2.Compile(match, re2.IgnoreCase)
			if err != nil {
				Log.Warnf(_f, "re2 compile %#q: %s", match, err)
				continue
			}

			r.match = exp
		}

		Log.Debugf(_f, "%#v", r.match)

		// can't error - already compiled
		m, _ := r.match.FindStringMatch(line)
		if m == nil || m.Index > 0 {
			continue
		}

		if r.Okay == nil {
			r.Okay = True
		}

		if r.Okay(c) {
			c.Route = r
			line = line[m.Index:]

			break
		}
	}

	cont := strings.TrimSpace(strings.TrimPrefix(line, pfx))
	args := Split(cont)

	if len(args) == 0 {
		return nil
	}

	cmd := args[0]
	args = args[1:]

	Log.Debugf(_f, "cmd %q, args %q", cmd, args)

	c.Flags = NewFlags(c, cmd, args)

	Log.Debugf(_f, "ctx %#v", c)

	if c.Route == nil {
		return nil
	}

	return c
}
