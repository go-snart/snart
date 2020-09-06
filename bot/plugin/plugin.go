// Package plugin provides an interface for loading functionality into Snart.
package plugin

import (
	"fmt"
	"os"
	"path/filepath"
	"plugin"

	dg "github.com/bwmarrin/discordgo"

	"github.com/go-snart/snart/bot/gamer"
	"github.com/go-snart/snart/db"
	"github.com/go-snart/snart/log"
	"github.com/go-snart/snart/route"
)

// PluginsEnv is the environment variable used to add plugin directories.
const PluginsEnv = "SNART_PLUGINS"

// Plugin provides all the components of a snart plugin.
type Plugin interface {
	Session(*dg.Session)
	DB(*db.DB)

	Intents() dg.Intent
	Routes() []*route.Route
	Gamers() []gamer.Gamer

	String() string
}

// OpenAll returns all of the Plugins.
func OpenAll() []Plugin {
	dirs := []string{"plugins"}

	env, ok := os.LookupEnv(PluginsEnv)
	if ok {
		dirs = append(dirs, filepath.SplitList(env)...)
	}

	plugs := []Plugin(nil)

	for _, dir := range dirs {
		dplugs := []Plugin(nil)

		glob := filepath.Join(dir, "*")
		log.Info.Printf("loading plugins from %q\n", glob)

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

			const name = "Plugin"

			sym, err := plug.Lookup(name)
			if err != nil {
				err = fmt.Errorf("lookup %q: %w", name, err)
				log.Warn.Println(err)

				continue
			}

			dplugs = append(dplugs, sym.(Plugin))
		}

		log.Info.Printf("loaded %d plugins - %v\n", len(dplugs), dplugs)

		plugs = append(plugs, dplugs...)
	}

	return plugs
}
