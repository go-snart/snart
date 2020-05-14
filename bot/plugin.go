package bot

var Plugins = make(map[string]Plugin)

type Plugin func(*Bot) error

func Register(name string, plug Plugin) {
	_f := "Register"
	if _, ok := Plugins[name]; ok {
		Log.Warnf(_f, "plugin %s already registered", name)
		return
	}
	Plugins[name] = plug
}

func warnPlug(_f string, b *Bot, name string, plug Plugin) {
	err := plug(b)
	if err != nil {
		Log.Warnf(_f, "plugin %s: %s", name, err)
	}
}

func (b *Bot) GoPlugins() {
	_f := "(*Bot).GoPlugins"
	for name, plug := range Plugins {
		go warnPlug(_f, b, name, plug)
	}
}
