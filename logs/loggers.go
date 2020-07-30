package logs

import (
	"os"

	"github.com/superloach/nilog"
)

func Debug(name string) *nilog.Logger {
	return nilog.New(os.Stdout, "(debug) "+name+": ", nilog.LstdFlags)
}

func Info(name string) *nilog.Logger {
	return nilog.New(os.Stdout, name+": ", nilog.LstdFlags)
}

func Warn(name string) *nilog.Logger {
	return nilog.New(os.Stderr, "(warn) "+name+": ", nilog.LstdFlags)
}

func Loggers(name string) (
	debug *nilog.Logger,
	info *nilog.Logger,
	warn *nilog.Logger,
) {
	return Debug(name),
		Info(name),
		Warn(name)
}
