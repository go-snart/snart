package token

import (
	"context"
	"fmt"

	"github.com/go-snart/snart/db"
)

func table(ctx context.Context, d *db.DB) {
	const (
		_f = "table"
		e  = `CREATE TABLE IF NOT EXISTS token(
			value TEXT
		)`
	)

	_, err := d.Conn(&ctx).Exec(ctx, e)
	if err != nil {
		err = fmt.Errorf("exec %#q: %w", e, err)

		Log.Error(_f, err)

		return
	}
}

// SelectTokens retrieves bot tokens from a DB.
func SelectTokens(ctx context.Context, d *db.DB) ([]string, error) {
	const _f = "DBToken"

	Log.Debug(_f, "enter")

	table(ctx, d)

	const q = `SELECT value FROM token`

	rows, err := d.Conn(&ctx).Query(ctx, q)
	if err != nil {
		err = fmt.Errorf("query %#q: %w", q, err)

		Log.Error(_f, err)

		return nil, err
	}
	defer rows.Close()

	toks := []string(nil)

	for rows.Next() {
		tok := ""

		err = rows.Scan(&tok)
		if err != nil {
			err = fmt.Errorf("scan tok: %w", err)

			Log.Error(_f, err)

			return nil, err
		}

		toks = append(toks, tok)
	}

	if err := rows.Err(); err != nil {
		err = fmt.Errorf("rows: %w", err)

		Log.Error(_f, err)

		return nil, err
	}

	return toks, nil
}

// InsertTokens adds tokens to the database so that they're persistent.
func InsertTokens(ctx context.Context, d *db.DB, toks []string) {
	const _f = "InsertTokens"

	table(ctx, d)

	e := `INSERT INTO token(value) VALUES`
	vals := []interface{}(nil)

	for n, tok := range toks {
		e += fmt.Sprintf(`($%d)`, n)

		if n < len(toks)-1 {
			e += `,`
		}

		vals = append(vals, tok)
	}

	_, err := d.Conn(&ctx).Exec(ctx, e, vals...)
	if err != nil {
		err = fmt.Errorf("exec %#q (%#v): %w", e, vals, err)

		Log.Warn(_f, err)

		return
	}
}
