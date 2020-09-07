package db

import (
	"context"
	"fmt"

	"github.com/go-snart/snart/log"
)

// Admins returns a list of known admin IDs from the database.
func (d *DB) Admins(ctx context.Context) ([]string, error) {
	admins, err := d.LRange(ctx, "admins", 0, -1).Result()
	if err != nil {
		err = fmt.Errorf("lrange admins 0 -1: %w", err)
		log.Warn.Println(err)

		return nil, err
	}

	return admins, nil
}
