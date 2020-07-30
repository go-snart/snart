package logs

import (
	"fmt"
	"os"
	"strconv"

	"github.com/superloach/nilog"
)

func Debug(name string) *nilog.Logger {
	const EnvName = "SNART_DEBUG"

	debug, err := strconv.ParseBool(os.Getenv(EnvName))
	if err != nil {
		err = fmt.Errorf("parse $%s: %w", EnvName, err)
		panic(err)
	}

	if debug {
		return nilog.New(os.Stdout, "(debug) "+name+": ", nilog.LstdFlags)
	}

	return nil
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
