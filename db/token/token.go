// Package token provides token stuff for db.
package token

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-snart/snart/db"
	"github.com/go-snart/snart/logs"
)

// Tokens returns a list of suitable tokens.
func Tokens(ctx context.Context, d *db.DB) []string {
	logs.Debug.Println("enter->env")

	allToks := []string(nil)

	toks, err := EnvTokens()
	if err != nil && !errors.Is(err, ErrEnvUnset) {
		err = fmt.Errorf("env tok: %w", err)
		logs.Warn.Println(err)
	} else {
		allToks = append(allToks, toks...)
	}

	logs.Debug.Println("env->select")

	toks, err = SelectTokens(ctx, d)
	if err != nil {
		err = fmt.Errorf("select tok: %w", err)
		logs.Warn.Println(err)
	} else {
		allToks = append(allToks, toks...)
	}

	logs.Debug.Println("select->stdin")

	if len(allToks) == 0 {
		toks, err = StdinTokens()
		if err != nil {
			err = fmt.Errorf("stdin tok: %w", err)
			logs.Warn.Println(err)
		} else {
			logs.Debug.Println("stdin->insert")

			InsertTokens(ctx, d, toks)

			logs.Debug.Println("insert->exit")

			allToks = append(allToks, toks...)
		}
	}

	logs.Debug.Println("stdin->exit")

	return allToks
}
