// Package route contains a command router for a Snart Bot.
package route

import (
	re2 "github.com/dlclark/regexp2"
	"github.com/superloach/minori"
)

// Log is the logger for the route package.
var Log = minori.GetLogger("route")

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
