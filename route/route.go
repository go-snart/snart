// Package route contains a command router for a Snart Bot.
package route

import (
	"fmt"

	re2 "github.com/dlclark/regexp2"

	"github.com/go-snart/snart/log"
)

// Route is a command route.
type Route struct {
	Name  string
	Match *re2.Regexp
	Cat   string
	Desc  string
	Okay  Okay
	Func  func(*Ctx) error
}

// MustMatch compiles a *re2.Regexp with sensible options for a Route.
func MustMatch(match string) *re2.Regexp {
	r, err := re2.Compile(match, re2.IgnoreCase)
	if err != nil {
		err = fmt.Errorf("re2 compile %q: %w", match, err)
		log.Warn.Fatalln(err)

		return nil
	}

	return r
}
