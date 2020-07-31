// Package logs provides utilities for getting sensible loggers.
package logs

import (
	"log"
	"os"

	"github.com/logrusorgru/aurora"
	"github.com/mattn/go-colorable"
	"github.com/superloach/nilog"
)

func info(name string) *nilog.Logger {
	return nilog.New(
		colorable.NewColorableStdout(),
		aurora.Green("[info ] "+name+": ").String(),
		log.LstdFlags,
	)
}

func warn(name string) *nilog.Logger {
	return nilog.New(
		colorable.NewColorableStderr(),
		aurora.Red("[warn ] "+name+": ").String(),
		log.LstdFlags,
	)

	return nilog.New(os.Stderr, "(warn) "+name+": ", log.LstdFlags)
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
