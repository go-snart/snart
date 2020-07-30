// Package token provides token stuff for db.
package token

import (
	"context"
	"fmt"

	"github.com/go-snart/snart/db"
)

// Log is the logger for token.
var Log = db.Log.GetLogger("token")

// Tokens returns a list of suitable tokens.
func Tokens(ctx context.Context, d *db.DB) []string {
	const _f = "Tokens"

	allToks := []string(nil)

	toks, err := EnvTokens()
	if err != nil {
		err = fmt.Errorf("env tok: %w", err)
		Log.Warn(_f, err)
	} else {
		allToks = append(allToks, toks...)
	}

	toks, err = SelectTokens(ctx, d)
	if err != nil {
		err = fmt.Errorf("select tok: %w", err)
		Log.Warn(_f, err)
	} else {
		allToks = append(allToks, toks...)
	}

	if len(allToks) == 0 {
		toks, err = StdinTokens()
		if err != nil {
			err = fmt.Errorf("stdin tok: %w", err)
			Log.Warn(_f, err)
		} else {
			InsertTokens(ctx, d, toks)

			allToks = append(allToks, toks...)
		}
	}

	return allToks
}
