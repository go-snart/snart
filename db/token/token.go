// Package token provides token stuff for db.
package token

import "github.com/go-snart/snart/db"

// Tokens returns a list of suitable tokens.
func Tokens(d *db.DB) []string {
	return append(
		db.EnvStrings(d.Name, "token"),
		GetTokens(d)...,
	)
}
