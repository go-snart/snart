package bot

import (
	"sync"

	"github.com/go-snart/snart/logs"
)

// Plugin is a function which registers a plugin onto a Bot.
type Plugin func(*Bot) error

var plugins = struct {
	sync.Mutex
	m map[string]Plugin
}{
	m: map[string]Plugin{},
}

// RegisterPlugin adds a Plugin to a list of plugins to be loaded at startup.
func RegisterPlugin(name string, plug Plugin) {
	plugins.Lock()
	defer plugins.Unlock()

	if _, ok := plugins.m[name]; ok {
		logs.Warn.Fatalf("attempted to register plugin %q twice", name)
	}

	plugins.m[name] = plug

	logs.Info.Printf("registered plugin %q", name)
}

func (b *Bot) goPlugins() {
	plugins.Lock()
	defer plugins.Unlock()

	for name, plug := range plugins.m {
		go func(n string, p Plugin) {
			err := p(b)
			if err != nil {
				logs.Info.Printf("plugin %s: %s", n, err)
				return
			}
		}(name, plug)
	}
}
