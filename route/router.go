package route

import (
	"errors"
	"fmt"
	"strings"

	dg "github.com/bwmarrin/discordgo"

	"github.com/go-snart/snart/db"
	"github.com/go-snart/snart/db/prefix"
	"github.com/go-snart/snart/logs"
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
func (rr *Router) Ctx(pfx *prefix.Prefix, s *dg.Session, m *dg.Message, line string) *Ctx {
	c := &Ctx{
		Prefix:  pfx,
		Session: s,
		Message: m,
		Flag:    nil,
		Route:   nil,
	}

	line = strings.TrimSpace(strings.TrimPrefix(line, pfx.Value))

	for _, r := range *rr {
		m, _ := r.Match.FindStringMatch(line)
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

	cont := strings.TrimSpace(strings.TrimPrefix(line, pfx.Value))
	args := Split(cont)

	if len(args) == 0 {
		return nil
	}

	cmd := args[0]
	args = args[1:]
	c.Flag = NewFlag(c, cmd, args)

	if c.Route == nil {
		return nil
	}

	return c
}

// Handler returns a discordgo handler function for the router.
func (rr *Router) Handler(d *db.DB) func(s *dg.Session, m *dg.MessageCreate) {
	return func(s *dg.Session, m *dg.MessageCreate) {
		logs.Debug.Println("handling")

		if m.Message.Author.ID == s.State.User.ID {
			logs.Debug.Println("ignore self")
			return
		}

		if m.Message.Author.Bot {
			logs.Debug.Println("ignore bot")
			return
		}

		lines := strings.Split(m.Message.Content, "\n")
		logs.Debug.Printf("lines %#v", lines)

		for _, line := range lines {
			logs.Debug.Printf("line %q", line)

			pfx, err := prefix.FindPrefix(d, s, m.GuildID, line)
			if err != nil {
				if errors.Is(err, prefix.ErrPrefixFail) {
					continue
				}

				err = fmt.Errorf("prefix %q %q: %w", m.GuildID, line, err)
				logs.Warn.Println(err)

				continue
			}

			c := rr.Ctx(pfx, s, m.Message, line)
			if c == nil {
				continue
			}

			err = c.Run()
			if err != nil {
				err = fmt.Errorf("c run: %w", err)
				logs.Warn.Println(err)

				continue
			}
		}
	}
}
