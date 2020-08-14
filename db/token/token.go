// Package token provides token stuff for db.
package token

import (
	"github.com/go-snart/snart/db"
)

// Tokens returns a list of suitable tokens.
func Tokens(d *db.DB) []string {
	stdinToks := db.StdinStrings("discord token")

	StoreTokens(d, stdinToks)

	return append(
		append(
			EnvTokens(),
			GetTokens(d)...,
		),
		stdinToks...,
	)
}
