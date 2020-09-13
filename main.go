// Package main is the command for Snart.
//
// The initialization and running are done under the same context.Background().
// The bot name is guessed from the name of the executable (args[0]).
package main

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-snart/bot"
)

func main() {
	ctx := context.Background()

	_, name := filepath.Split(os.Args[0])
	name = strings.Split(name, ".")[0]

	// nolint:errcheck
	bot.New(ctx, name).Run(ctx)
}
