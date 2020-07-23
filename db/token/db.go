// Package token provides auth token stuff for db.
package token

import (
	"context"
	"fmt"

	"github.com/go-snart/snart/db"
)

// Log is the logger for token.
var Log = db.Log.GetLogger("token")

// Table is a table builder for config.admin.
func Table(ctx context.Context, d *db.DB) {
	x, err := d.Conn(&ctx).Exec(ctx, `CREATE TABLE IF NOT EXISTS token(value TEXT)`)
	Log.Debugf("Table", "%#v %#v", x, err)
}

// Token retrieves a token for a Bot.
func Token(ctx context.Context, d *db.DB) (string, error) {
	_f := "(*DB).Token"
	Log.Debug(_f, "enter")

	Table(ctx, d)

	const q = `SELECT value FROM token`

	rows, err := d.Conn(&ctx).Query(ctx, q)
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

	x, err := d.Conn(&ctx).Exec(ctx, q2, token)
	if err != nil {
		err = fmt.Errorf("db exec %#q(%q): %w", q2, token, err)
		Log.Error(_f, err)

		return "", err
	}

	Log.Info(_f, x)

	return token, nil
}
