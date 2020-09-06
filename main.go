// Package main provides an example bot using Snart.
package main

import (
	"os"
	"path/filepath"

	"github.com/go-snart/snart/bot"
)

func main() {
	_, name := filepath.Split(os.Args[0])
	bot.New(name).Run()
}
