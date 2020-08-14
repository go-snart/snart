package test

import "github.com/go-snart/snart/route"

// Flag gets a test *route.Flag.
func Flag(content string) *route.Flag {
	return Ctx(content).Flag
}
