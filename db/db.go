// Package db contains the workings of a DB for a Snart Bot.
package db

import (
	"fmt"

	redis "gopkg.in/redis.v5"

	"github.com/go-snart/snart/logs"
)

// DB wraps a redis.Cmdable, along with a Name and Opts used in creation.
type DB struct {
	redis.Cmdable

	Name string
	Opts *redis.Options
}

// New creates a DB using redis.NewClient.
func New(name string) *DB {
	for _, connString := range ConnStrings(name) {
		opts, err := redis.ParseURL(connString)
		if err != nil {
			err = fmt.Errorf("parse url %q: %w", connString, err)
			logs.Warn.Println(err)

			continue
		}

		c := redis.NewClient(opts)

		if c.Ping().Err() != nil {
			continue
		}

		return NewFromCmdable(c, name, opts)
	}

	logs.Warn.Fatalln("no good options found")

	return nil
}

// NewFromCmdable creates a DB using redis.NewClient.
func NewFromCmdable(c redis.Cmdable, name string, opts *redis.Options) *DB {
	return &DB{
		Cmdable: c,
		Name:    name,
		Opts:    opts,
	}
}
