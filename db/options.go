package db

import (
	"fmt"

	redis "gopkg.in/redis.v5"

	"github.com/go-snart/snart/logs"
)

// ConnStrings returns a list of useable redis conn strings.
func ConnStrings(name string) []string {
	allConnStrings := EnvConnStrings(name)

	if len(allConnStrings) == 0 {
		allConnStrings = append(allConnStrings, StdinConnStrings(name)...)
	}

	return allConnStrings
}

// Options returns a list of useable redis options.
func Options(name string) []*redis.Options {
	opts := []*redis.Options(nil)

	for _, connString := range ConnStrings(name) {
		opt, err := redis.ParseURL(connString)
		if err != nil {
			err = fmt.Errorf("parse url %q: %w", connString, err)

			logs.Warn.Println(err)

			continue
		}

		opts = append(opts, opt)
	}

	return opts
}
