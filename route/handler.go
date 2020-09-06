package route

import (
	"fmt"
	"strings"

	dg "github.com/bwmarrin/discordgo"

	"github.com/go-snart/snart/db"
	"github.com/go-snart/snart/db/prefix"
	"github.com/go-snart/snart/log"
)

// Handler is a slice of Routes.
type Handler []*Route

// Add adds a Route to the Handler.
func (h *Handler) Add(rs ...*Route) {
	*h = append(*h, rs...)
}

// Ctx gets a Ctx by finding an appropriate Route for a given prefix, session, message, etc.
func (h *Handler) Ctx(pfx *prefix.Prefix, s *dg.Session, m *dg.Message, line string) *Ctx {
	c := &Ctx{
		Prefix:  pfx,
		Session: s,
		Message: m,
		Flag:    nil,
		Route:   nil,
	}

	log.Debug.Println("line", line)

	line = strings.TrimSpace(strings.TrimPrefix(line, pfx.Value))

	log.Debug.Println("line", line)

	args := Split(line)

	log.Debug.Println("args", args)

	if len(args) == 0 {
		log.Debug.Println("0 args")

		return nil
	}

	cmd := args[0]
	log.Debug.Println("cmd", cmd)

	args = args[1:]
	log.Debug.Println("args", args)

	for _, r := range *h {
		m, _ := r.Match.FindStringMatch(cmd)
		log.Debug.Println("m", m)

		if m == nil || m.Index > 0 {
			continue
		}

		if r.Okay == nil {
			r.Okay = True
		}

		if r.Okay(c) {
			c.Route = r

			break
		}
	}

	log.Debug.Println("route", c.Route)

	if c.Route == nil {
		return nil
	}

	c.Flag = NewFlag(c, cmd, args)
	log.Debug.Println("flag", c.Flag)

	return c
}

// Handle returns a discordgo handler function for the Handler.
func (h *Handler) Handle(d *db.DB) func(s *dg.Session, m *dg.MessageCreate) {
	return func(s *dg.Session, m *dg.MessageCreate) {
		log.Debug.Println("handling")

		if m.Message.Author.ID == s.State.User.ID {
			log.Debug.Println("ignore self")

			return
		}

		if m.Message.Author.Bot {
			log.Debug.Println("ignore bot")

			return
		}

		lines := strings.Split(m.Message.Content, "\n")
		log.Debug.Printf("lines %#v", lines)

		for _, line := range lines {
			log.Debug.Printf("line %q", line)

			pfx, err := prefix.FindPrefix(d, s, m.GuildID, line)
			if err != nil {
				err = fmt.Errorf("prefix %q %q: %w", m.GuildID, line, err)
				log.Warn.Println(err)

				continue
			}

			if pfx == nil {
				log.Warn.Println("nil pfx")

				continue
			}

			c := h.Ctx(pfx, s, m.Message, line)
			if c == nil {
				continue
			}

			err = c.Run()
			if err != nil {
				err = fmt.Errorf("c run: %w", err)
				log.Warn.Println(err)

				continue
			}
		}
	}
}
