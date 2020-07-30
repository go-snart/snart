// Package route contains a command router for a Snart Bot.
package route

import (
	re2 "github.com/dlclark/regexp2"

	"github.com/go-snart/snart/logs"
)

const _p = "route"

// Log is the logger for the route package.
var Debug, Info, Warn = logs.Loggers(_p)

// Route is a command route.
type Route struct {
	Name  string
	Match string
	match *re2.Regexp
	Cat   string
	Desc  string
	Okay  Okay
	Func  func(*Ctx) error
}
