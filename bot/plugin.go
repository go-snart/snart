package bot

// Plugins holds the plugins to be loaded into a Bot on startup.
var Plugins = make(map[string]Plugin)

// Plugin is a function which registers a plugin onto a Bot.
type Plugin func(*Bot) error

// Register adds a Plugin to the Plugins.
func Register(name string, plug Plugin) {
	_f := "Register"

	if _, ok := Plugins[name]; ok {
		Log.Fatalf(_f, "attempted to register plugin %q twice", name)
	}

	Plugins[name] = plug

	Log.Infof(_f, "registered plugin %q", name)
}

// GoPlugins spawns all of the Plugins on the Bot.
func (b *Bot) GoPlugins() {
	_f := "(*Bot).GoPlugins"

	for name, plug := range Plugins {
		go func(n string, p Plugin) {
			err := p(b)
			if err != nil {
				Log.Warnf(_f, "plugin %s: %s", n, err)
				return
			}
		}(name, plug)
	}
}
