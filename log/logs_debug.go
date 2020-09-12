// +build debug

package log

import (
	"log"

	"github.com/logrusorgru/aurora"
	"github.com/mattn/go-colorable"
	"github.com/superloach/nilog"
)

const flags = log.Llongfile

// Debug is a *nilog.Logger for debugging purposes, and is nil unless the debug tag is enabled.
var Debug = nilog.New(
	colorable.NewColorableStderr(),
	aurora.Yellow("[debug] ").String(),
	flags,
)
