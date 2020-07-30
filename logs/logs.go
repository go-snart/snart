// Package logs provides utilities for getting sensible loggers.
package logs

import (
	"fmt"
	"os"
	"strconv"

	"github.com/superloach/nilog"
)

func debug(name string) *nilog.Logger {
	const EnvName = "SNART_DEBUG"

	env, ok := os.LookupEnv(EnvName)
	if !ok {
		return nil
	}

	debug, err := strconv.ParseBool(env)
	if err != nil {
		err = fmt.Errorf("parse $%s: %w", EnvName, err)
		panic(err)
	}

	if debug {
		return nilog.New(os.Stdout, "(debug) "+name+": ", nilog.LstdFlags)
	}

	return nil
}

func info(name string) *nilog.Logger {
	return nilog.New(os.Stdout, name+": ", nilog.LstdFlags)
}

func warn(name string) *nilog.Logger {
	return nilog.New(os.Stderr, "(warn) "+name+": ", nilog.LstdFlags)
}

// Loggers returns sensible loggers for a package with the given name.
func Loggers(name string) (
	*nilog.Logger,
	*nilog.Logger,
	*nilog.Logger,
) {
	return debug(name),
		info(name),
		warn(name)
}
