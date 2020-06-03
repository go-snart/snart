package bot

import (
	"fmt"
	"os"
	"time"

	"github.com/go-snart/snart/db"
	"github.com/go-snart/snart/route"

	dg "github.com/bwmarrin/discordgo"
)

type Bot struct {
	DB      *db.DB
	Session *dg.Session

	Router *route.Router
	Gamers []Gamer

	Sig     chan os.Signal
	Startup time.Time
}

func MkBot(d *db.DB) (*Bot, error) {
	_f := "MkBot"
	b := &Bot{}

	Log.Info(_f, "making bot")

	b.DB = d

	s, err := dg.New()
	if err != nil {
		err = fmt.Errorf("dg new: %w", err)
		Log.Error(_f, err)
		return nil, err
	}
	b.Session = s
	b.Session.AddHandler(b.Route)

	b.Router = route.NewRouter()
	b.Gamers = []Gamer{
		GamerUptime,
	}

	b.Sig = make(chan os.Signal)

	Log.Info(_f, "made bot")
	return b, nil
}
