package bot

// Plugins holds the plugins to be loaded into a Bot on startup.
var Plugins = make(map[string]Plugin)

// Plugin is a function which registers a plugin onto a Bot.
type Plugin func(*Bot) error

// Register adds a Plugin to the Plugins.
func Register(name string, plug Plugin) {
	_f := "Register"

	if _, ok := Plugins[name]; ok {
		Log.Warnf(_f, "plugin %s already registered", name)
		return
	}

	Plugins[name] = plug

	Log.Infof(_f, "registered as %s: %#v", name, plug)
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

			Log.Infof(_f, "plugin %s :)", n)
		}(name, plug)
	}
}
