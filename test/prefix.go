package test

import "github.com/go-snart/snart/db/prefix"

const (
	// PrefixValue is the prefix value used by Prefix.
	PrefixValue = "./"

	// PrefixClean is the clean prefix value used by Prefix.
	PrefixClean = "./"
)

// Prefix gets a test *prefix.Prefix.
func Prefix() *prefix.Prefix {
	return &prefix.Prefix{
		Value: PrefixValue,
		Clean: PrefixClean,
	}
}
