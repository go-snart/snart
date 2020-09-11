package db

import (
	"context"
	"fmt"

	"github.com/gomodule/redigo/redis"

	"github.com/go-snart/snart/log"
)

// Admins returns a list of known admin IDs from the database.
func (d *DB) Admins(ctx context.Context) ([]string, error) {
	conn, err := d.GetContext(ctx)
	if err != nil {
		err = fmt.Errorf("get context: %w", err)
		log.Warn.Println(err)

		return nil, err
	}
	defer conn.Close()

	admins, err := redis.Strings(conn.Do("LRANGE", "admins", 0, -1))
	if err != nil {
		err = fmt.Errorf("lrange admins: %w", err)
		log.Warn.Println(err)

		return nil, err
	}

	return admins, nil
}
