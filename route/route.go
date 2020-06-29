package route

import (
	re2 "github.com/dlclark/regexp2"
	"github.com/superloach/minori"
)

var Log = minori.GetLogger("route")

type Route struct {
	Name  string
	Match string
	match *re2.Regexp
	Cat   string
	Desc  string
	Okay  Okay
	Func  func(*Ctx) error
}
