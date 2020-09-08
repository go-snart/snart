package db

import (
	"context"
	"fmt"

	"github.com/go-snart/snart/log"
)

// Tokens returns a list of suitable tokens.
func (d *DB) Tokens(ctx context.Context) []string {
	toks := EnvStrings(d.Name, "token")

	dtoks, err := d.LRange(ctx, "tokens", 0, -1).Result()
	if err != nil {
		err = fmt.Errorf("lrange tokens 0 -1: %w", err)
		log.Warn.Println(err)

		return nil
	}

	return append(toks, dtoks...)
}
