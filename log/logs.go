// Package log provides utilities for getting sensible loggers.
package log

import (
	"github.com/logrusorgru/aurora"
	"github.com/mattn/go-colorable"
	"github.com/superloach/nilog"
)

func init() {
	Debug.Println("debugging :)")
}

// Info is a *nilog.Logger for healthy, informational logs.
var Info = nilog.New(
	colorable.NewColorableStdout(),
	aurora.Green("[info ] ").String(),
	flags,
)

// Warn is a *nilog.Logger for unhealthy, warning logs.
var Warn = nilog.New(
	colorable.NewColorableStderr(),
	aurora.Red("[warn ] ").String(),
	flags,
)
