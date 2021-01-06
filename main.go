// Package main is an example command for Snart.
package main

import (
	"log"
	"os"

	"github.com/go-snart/snart/admin"
	"github.com/go-snart/snart/bot"
)

func main() {
	db := os.Getenv("SNART_DB")

	b, err := bot.Open(db)
	if err != nil {
		log.Fatalf("open: %s", err)
	}

	err = b.Plug(admin.Plug)
	if err != nil {
		log.Fatalf("plug admin: %s", err)
	}

	err = b.Run()
	if err != nil {
		log.Fatalf("run: %s", err)
	}
}
