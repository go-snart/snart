package route

import (
	re "regexp"

	"github.com/superloach/minori"
)

var Log = minori.GetLogger("route")

type Route struct {
	Name  string
	Match string
	match *re.Regexp
	Cat   string
	Desc  string
	Okay  Okay
	Func  func(*Ctx) error
}
