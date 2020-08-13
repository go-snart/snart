// Package token provides token stuff for db.
package token

import (
	"fmt"

	"github.com/go-snart/snart/db"
	"github.com/go-snart/snart/logs"
)

// Tokens returns a list of suitable tokens.
func Tokens(d *db.DB) []string {
	logs.Debug.Println("enter->env")

	allToks := EnvTokens()

	logs.Debug.Println("env->select")

	toks, err := GetTokens(d)
	if err != nil {
		err = fmt.Errorf("select tok: %w", err)
		logs.Warn.Println(err)
	} else {
		allToks = append(allToks, toks...)
	}

	logs.Debug.Println("select->stdin")

	if len(allToks) == 0 {
		toks = StdinTokens()

		allToks = append(allToks, toks...)

		logs.Debug.Println("stdin->insert")

		err := StoreTokens(d, toks)
		if err != nil {
			err = fmt.Errorf("store tokens %v: %w", toks, err)
			logs.Warn.Println(err)
			return allToks
		}

		logs.Debug.Println("insert->exit")
	}

	logs.Debug.Println("stdin->exit")

	return allToks
}
