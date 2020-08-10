// Package route contains a command router for a Snart Bot.
package route

import re2 "github.com/dlclark/regexp2"

// Route is a command route.
type Route struct {
	Name string

	Match string
	match *re2.Regexp

	Cat  string
	Desc string
	Okay Okay
	Func func(*Ctx) error
}
