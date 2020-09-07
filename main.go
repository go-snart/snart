// Package main provides an example bot using Snart.
package main

import (
	"context"
	"os"
	"path/filepath"

	"github.com/go-snart/snart/bot"
)

func main() {
	ctx := context.Background()
	_, name := filepath.Split(os.Args[0])

	bot.New(ctx, name).Run(ctx)
}
