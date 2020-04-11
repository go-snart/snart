package main

import (
	"flag"
	"fmt"

	"github.com/superloach/minori"

	"github.com/go-snart/bot"
)

var debug = flag.Bool("debug", false, "print debug messages")
var plugins = flag.String("plugins", "", "`dir` to load plugins from")
var dburl = flag.String("dburl", ":28015", "rethinkdb url")

var Log = minori.GetLogger("cmd/snart")

func main() {
	_f := "main"
	flag.Parse()

	if *dburl == "" {
		Log.Fatal(_f, "empty dburl")
	}

	if *debug {
		minori.Level = minori.DEBUG
		Log.Debug(_f, "debug mode")
	} else {
		minori.Level = minori.INFO
	}

	// make bot
	b, err := bot.MkBot(*dburl)
	if err != nil {
		err = fmt.Errorf("mkbot %#v: %w", *dburl, err)
		Log.Fatal(_f, err)
	}

	if *plugins == "" {
		Log.Warn(_f, "trying to load plugins from .")
		*plugins = "."
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
