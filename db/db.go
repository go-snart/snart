// Package db contains the workings of a DB for a Snart Bot.
package db

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"

	"github.com/go-snart/snart/log"
)

// DB wraps a redis.Cmdable, along with a Name and Opts used in creation.
type DB struct {
	redis.Cmdable

	Name string
	Opts *redis.Options
}

// New creates a DB using redis.NewClient.
func New(ctx context.Context, name string) *DB {
	for _, connString := range append(
		EnvStrings(name, "db"),

		"redis://"+name+"_db",
		"redis://"+name+"_db.docker",
	) {
		opts, err := redis.ParseURL(connString)
		if err != nil {
			err = fmt.Errorf("parse url %q: %w", connString, err)
			log.Warn.Println(err)

			continue
		}

		log.Debug.Printf("trying %q\n", connString)

		c := redis.NewClient(opts)
		if c.Ping(ctx).Err() != nil {
			log.Debug.Println("ping failed")

			continue
		}

		return NewFromCmdable(c, name, opts)
	}

	log.Warn.Fatalf("no good options found for db %q", name)

	return nil
}

// NewFromCmdable creates a DB using redis.NewClient.
func NewFromCmdable(c redis.Cmdable, name string, opts *redis.Options) *DB {
	return &DB{
		Cmdable: c,

		Name: name,
		Opts: opts,
	}
}
