package route

import "github.com/superloach/minori"

var Log = minori.GetLogger("route")

type Route struct {
	Name  string
	Match string
	Cat   string
	Desc  string
	Okay  Okay
	Func  func(*Ctx) error
}
