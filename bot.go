package bot

import (
	"fmt"
	"os"
	"time"

	"github.com/go-snart/db"
	"github.com/go-snart/route"
	"github.com/superloach/minori"

	dg "github.com/bwmarrin/discordgo"
)

type Bot struct {
	Startup time.Time
	Session *dg.Session
	DB      *db.DB
	Routes  []*route.Route
	Sig     chan os.Signal
}

var Log = minori.GetLogger("bot")

func MkBot(dburl string) (*Bot, error) {
	_f := "MkBot"
	b := &Bot{}

	Log.Info(_f, "making bot")

	s, err := dg.New()
	if err != nil {
		err = fmt.Errorf("dg new: %w", err)
		Log.Error(_f, err)
		return nil, err
	}
	b.Session = s
	b.Session.AddHandler(b.Route)

	b.DB = &db.DB{URL: dburl}

	b.Routes = make([]*route.Route, 0)
	b.Sig = make(chan os.Signal)

	Log.Info(_f, "made bot")
	return b, nil
}
