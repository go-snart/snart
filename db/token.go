package db

import (
	"errors"
	"fmt"
)

// ErrTokenFail indicates a function failed to get a token.
var ErrTokenFail = errors.New("failed to get a token")

// Token contains the Value of a Discord bot token.
type Token struct {
	Value string
}

// TokenTable is a table builder for config.token.
var TokenTable = BuildTable(
	ConfigDB, "token",
	nil, nil,
)

// Token retrieves a token for a Bot.
func (d *DB) Token() (*Token, error) {
	_f := "(*DB).Token"
	Log.Debug(_f, "enter")

	toks := make([]*Token, 0)

	err := TokenTable.ReadAll(&toks, d)
	if err != nil {
		err = fmt.Errorf("readall &toks: %w", err)
		Log.Error(_f, err)

		return nil, err
	}

	if len(toks) < 1 {
		return nil, ErrTokenFail
	}

	return toks[0], nil
}
