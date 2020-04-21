package main

import (
	"fmt"

	"github.com/go-snart/bot"
	"github.com/namsral/flag"
	"github.com/superloach/minori"
)

var (
	debug = flag.Bool("debug", false, "print debug messages")

	dburl  = flag.String("dburl", "", "rethinkdb url")
	dbuser = flag.String("dbuser", "", "rethinkdb username")
	dbpass = flag.String("dbpass", "", "rethinkdb password")

	plugins = flag.String("plugins", "", "dir to load plugins from")
)

var Log = minori.GetLogger("snart")

func main() {
	_f := "main"
	flag.Parse()

	if *debug {
		minori.Level = minori.DEBUG
		Log.Debug(_f, "debug mode")
	} else {
		minori.Level = minori.INFO
	}

	// make bot
	b, err := bot.MkBot(*dburl, *dbuser, *dbpass)
	if err != nil {
		err = fmt.Errorf("mkbot %#v: %w", *dburl, err)
		Log.Fatal(_f, err)
	}

	if *plugins == "" {
		Log.Warn(_f, "trying to load plugins from ./plugins")
		*plugins = "./plugins"
	}

	// register plugins
	err = b.RegisterAll(*plugins)
	if err != nil {
		err = fmt.Errorf("registerall %#v: %w", *plugins, err)
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
