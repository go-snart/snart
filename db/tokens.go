package db

import (
	"context"
	"fmt"

	"github.com/gomodule/redigo/redis"

	"github.com/go-snart/snart/log"
)

// Tokens returns a list of suitable tokens.
func (d *DB) Tokens(ctx context.Context) ([]string, error) {
	toks := EnvStrings(d.Name, "token")

	conn, err := d.GetContext(ctx)
	if err != nil {
		err = fmt.Errorf("get conn: %w", err)
		log.Warn.Println(err)

		return nil, err
	}
	defer conn.Close()

	dtoks, err := redis.Strings(conn.Do("LRANGE", "tokens", 0, -1))
	if err != nil {
		err = fmt.Errorf("lrange tokens 0 -1: %w", err)
		log.Warn.Println(err)

		return nil, err
	}

	return append(toks, dtoks...), nil
}
