// +build !snart_debug

package logs

import (
	"log"

	"github.com/superloach/nilog"
)

const flags = log.Lshortfile

var Debug = (*nilog.Logger)(nil)
