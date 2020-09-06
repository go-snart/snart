// Package plug provides plugin functionality for Snart.
package plug

import (
	"fmt"
	"path/filepath"
	"plugin"

	dg "github.com/bwmarrin/discordgo"

	"github.com/go-snart/snart/bot/gamer"
	"github.com/go-snart/snart/bot/halt"
	"github.com/go-snart/snart/db"
	"github.com/go-snart/snart/log"
	"github.com/go-snart/snart/route"
)

// Plug provides all the components of a Snart plugin.
type Plug interface {
	String() string

	PlugDB(*db.DB)
	PlugSession(*dg.Session)
	PlugHalt(chan halt.Halt)

	Routes() []*route.Route
	Intents() dg.Intent
	Gamers() []gamer.Gamer
}

// Plugs loads and returns all of the Plugs.
func Plugs(name string) []Plug {
	dirs := append(
		[]string{"plugs"},
		db.EnvStrings(name, "plugs")...,
	)

	plugs := []Plug(nil)

	for _, dir := range dirs {
		dplugs := []Plug(nil)

		glob := filepath.Join(dir, "*")
		log.Info.Printf("loading plugs from %q\n", glob)

		pfs, err := filepath.Glob(glob)
		if err != nil {
			err = fmt.Errorf("glob %q: %w", glob, err)
			log.Warn.Println(err)

			continue
		}

		for _, pf := range pfs {
			plug, err := plugin.Open(pf)
			if err != nil {
				err = fmt.Errorf("plugin open %q: %w", pf, err)
				log.Warn.Println(err)

				continue
			}

			const name = "Plug"

			sym, err := plug.Lookup(name)
			if err != nil {
				err = fmt.Errorf("lookup %q: %w", name, err)
				log.Warn.Println(err)

				continue
			}

			psym, ok := sym.(Plug)
			if !ok {
				log.Warn.Println("sym was not a Plug")

				continue
			}

			dplugs = append(dplugs, psym)
		}

		log.Info.Printf("loaded %d plugs - %v\n", len(dplugs), dplugs)

		plugs = append(plugs, dplugs...)
	}

	return plugs
}
