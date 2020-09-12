// Package db contains the workings of a DB for a Snart Bot.
package db

import (
	"context"
	"time"

	"github.com/gomodule/redigo/redis"

	"github.com/go-snart/snart/log"
)

// DB wraps a redis.Cmdable, along with a Name and Opts used in creation.
type DB struct {
	Name string
	*redis.Pool
}

// New creates a DB using redis.NewClient.
func New(ctx context.Context, name string) *DB {
	for _, addr := range append(
		EnvStrings(name, "db"),

		// convenient for docker dns
		name+"_db:6379",
		name+"_db.docker:6379",
	) {
		addr := addr
		pool := &redis.Pool{
			MaxIdle:     3,
			IdleTimeout: 240 * time.Second,
			Wait:        true,
			DialContext: func(ctx context.Context) (redis.Conn, error) {
				return redis.DialContext(ctx, "tcp", addr)
			},
		}

		log.Debug.Printf("trying %q\n", addr)

		conn, err := pool.GetContext(ctx)
		if err != nil {
			log.Debug.Printf("get context failed: %s\n", err)

			continue
		}
		defer conn.Close()

		_, err = conn.Do("PING")
		if err != nil {
			log.Debug.Printf("ping failed: %s\n", err)

			continue
		}

		return &DB{
			Name: name,
			Pool: pool,
		}
	}

	log.Warn.Fatalf("no good options found for db %q", name)

	return nil
}
