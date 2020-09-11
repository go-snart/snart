package route

import (
	"context"
	"fmt"
	"strings"

	dg "github.com/bwmarrin/discordgo"

	"github.com/go-snart/snart/db"
	"github.com/go-snart/snart/log"
)

// Handler is a slice of Routes.
type Handler struct {
	Routes []*Route
	DB     *db.DB
}

// NewHandler creates a Handler.
func NewHandler(d *db.DB) *Handler {
	return &Handler{
		Routes: []*Route(nil),
		DB:     d,
	}
}

// Add adds a Route to the Handler.
func (h *Handler) Add(rs ...*Route) {
	h.Routes = append(h.Routes, rs...)
}

// Ctx gets a Ctx by finding an appropriate Route for a given prefix, session, message, etc.
func (h *Handler) Ctx(
	ctx context.Context,
	pfx *db.Prefix,
	s *dg.Session,
	m *dg.Message,
	line string,
) *Ctx {
	c := &Ctx{
		Context: ctx,
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

	for _, r := range h.Routes {
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
func (h *Handler) Handle(s *dg.Session, m *dg.MessageCreate) {
	ctx := context.Background()

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

		pfx, err := h.DB.FindPrefix(ctx, s, m.GuildID, line)
		if err != nil {
			err = fmt.Errorf("prefix %q %q: %w", m.GuildID, line, err)
			log.Warn.Println(err)

			continue
		}

		if pfx == nil {
			continue
		}

		c := h.Ctx(ctx, pfx, s, m.Message, line)
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
