// Package logs provides utilities for getting sensible loggers.
package logs

import (
	"github.com/logrusorgru/aurora"
	"github.com/mattn/go-colorable"
	"github.com/superloach/nilog"
)

var Info = nilog.New(
	colorable.NewColorableStdout(),
	aurora.Green("[info ] ").String(),
	flags,
)

var Warn = nilog.New(
	colorable.NewColorableStderr(),
	aurora.Red("[warn ] ").String(),
	flags,
)
