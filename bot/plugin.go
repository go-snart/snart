package bot

// Plugins holds the plugins to be loaded into a Bot on startup.
var Plugins = make(map[string]Plugin)

// Plugin is a function which registers a plugin onto a Bot.
type Plugin func(*Bot) error

// Register adds a Plugin to the Plugins.
func Register(name string, plug Plugin) {
	if _, ok := Plugins[name]; ok {
		Info.Fatalf("attempted to register plugin %q twice", name)
	}

	Plugins[name] = plug

	Info.Printf("registered plugin %q", name)
}

// GoPlugins spawns all of the Plugins on the Bot.
func (b *Bot) GoPlugins() {
	for name, plug := range Plugins {
		go func(n string, p Plugin) {
			err := p(b)
			if err != nil {
				Info.Printf("plugin %s: %s", n, err)
				return
			}
		}(name, plug)
	}
}
