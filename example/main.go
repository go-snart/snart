package main

import (
	"fmt"

	"github.com/namsral/flag"
	"github.com/superloach/minori"

	"github.com/go-snart/snart/bot"
	"github.com/go-snart/snart/db"

	_ "github.com/go-snart/plugin-admin"
	_ "github.com/go-snart/plugin-help"
)

var (
	debug = flag.Bool("debug", false, "print debug messages")

	dbhost = flag.String("dbhost", "localhost", "rethinkdb host")
	dbport = flag.Int("dbport", 28015, "rethinkdb port")
	dbuser = flag.String("dbuser", "admin", "rethinkdb username")
	dbpass = flag.String("dbpass", "", "rethinkdb password")
)

var Log = minori.GetLogger("example")

func main() {
	_f := "main"
	flag.Parse()

	if *debug {
		minori.Level = minori.DEBUG
		Log.Debug(_f, "debug mode")
	} else {
		minori.Level = minori.INFO
	}

	Log.Debugf(_f, "plugins: %v", bot.Plugins)

	d := &db.DB{
		Host: *dbhost,
		Port: *dbport,
		User: *dbuser,
		Pass: *dbpass,
	}

	// make bot
	b, err := bot.NewBot(d)
	if err != nil {
		err = fmt.Errorf("new bot %#v: %w", d, err)
		Log.Fatal(_f, err)
	}

	// run the bot
	err = b.Start()
	if err != nil {
		err = fmt.Errorf("start: %w", err)
		Log.Fatal(_f, err)
	}

	Log.Info(_f, "bye!")
}
