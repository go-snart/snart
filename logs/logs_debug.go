// +build snart_debug

package logs

import (
	"log"

	"github.com/logrusorgru/aurora"
	"github.com/mattn/go-colorable"
	"github.com/superloach/nilog"
)

const flags = log.Llongfile

var Debug = nilog.New(
	colorable.NewColorableStderr(),
	aurora.Yellow("[debug] ").String(),
	flags,
)
