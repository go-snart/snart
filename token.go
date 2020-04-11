package bot

import (
	"errors"

	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

var TokenFail = errors.New("failed to get a token")

type Token struct {
	Value string
}

var TokenTable = r.DB("config").TableCreate("token")

func (b *Bot) Token() (*Token, error) {
	_f := "(*Bot).Token"
	Log.Debug(_f, "enter")

	b.DB.Easy(ConfigDB)
	b.DB.Easy(TokenTable)

	toks := make([]*Token, 0)
	q := r.DB("config").Table("token")
	err := q.ReadAll(&toks, b.DB)
	if err != nil {
		err = fmt.Errorf("readall &toks: %w", err)
		Log.Error(_f, err)
		return nil, err
	}

	if len(toks) == 0 {
		return nil, TokenFail
	}

	return toks[0], nil
}
