// Package logs provides utilities for getting sensible loggers.
package logs

import (
	"log"

	"github.com/logrusorgru/aurora"
	"github.com/mattn/go-colorable"
	"github.com/superloach/nilog"
)

const flags = log.Lshortfile

func info(name string) *nilog.Logger {
	return nilog.New(
		colorable.NewColorableStdout(),
		aurora.Green("[info ] ").String()+name+":",
		flags,
	)
}

func warn(name string) *nilog.Logger {
	return nilog.New(
		colorable.NewColorableStderr(),
		aurora.Red("[warn ] ").String()+name+":",
		flags,
	)
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
