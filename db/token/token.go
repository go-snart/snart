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

	Log.Debug(_f, "enter->env")

	allToks := []string(nil)

	toks, err := EnvTokens()
	if err != nil {
		err = fmt.Errorf("env tok: %w", err)
		Log.Warn(_f, err)
	} else {
		allToks = append(allToks, toks...)
	}

	Log.Debug(_f, "env->select")

	toks, err = SelectTokens(ctx, d)
	if err != nil {
		err = fmt.Errorf("select tok: %w", err)
		Log.Warn(_f, err)
	} else {
		allToks = append(allToks, toks...)
	}

	Log.Debug(_f, "select->stdin")

	if len(allToks) == 0 {
		toks, err = StdinTokens()
		if err != nil {
			err = fmt.Errorf("stdin tok: %w", err)
			Log.Warn(_f, err)
		} else {
			Log.Debug(_f, "stdin->insert")

			InsertTokens(ctx, d, toks)

			Log.Debug(_f, "insert->exit")

			allToks = append(allToks, toks...)
		}
	}

	Log.Debug(_f, "stdin->exit")

	return allToks
}
