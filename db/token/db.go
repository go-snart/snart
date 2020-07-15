package token

import (
	"context"
	"fmt"

	"github.com/go-snart/snart/db"
)

var Log = db.Log.GetLogger("token")

// TokenTable is a table builder for config.admin.
func TokenTable(ctx context.Context, d *db.DB) {
	x, err := d.Exec(ctx, `CREATE TABLE IF NOT EXISTS token(value TEXT)`)
	Log.Debugf("Tokentable", "%#v %#v", x, err)
}

// Token retrieves a token for a Bot.
func Token(ctx context.Context, d *db.DB) (string, error) {
	_f := "(*DB).Token"
	Log.Debug(_f, "enter")

	TokenTable(ctx, d)

	const q = `SELECT (value) FROM token`

	rows, err := d.Query(ctx, q)
	if err != nil {
		err = fmt.Errorf("db query %#q: %w", q, err)
		Log.Error(_f, err)

		return "", err
	}
	defer rows.Close()

	if rows.Next() {
		token := ""

		err = rows.Scan(&token)
		if err != nil {
			err = fmt.Errorf("scan token: %w", err)
			Log.Error(_f, err)

			return "", err
		}

		return token, nil
	}

	token, err := ScanToken()
	if err != nil {
		err = fmt.Errorf("scan token (cli): %w", err)
		Log.Error(_f, err)

		return "", err
	}

	const q2 = `INSERT INTO token(value) VALUES($1);`

	x, err := d.Exec(ctx, q2, token)
	if err != nil {
		err = fmt.Errorf("db exec %#q(%q): %w", q2, token, err)
		Log.Error(_f, err)

		return "", err
	}

	Log.Info(_f, x)

	return token, nil
}
