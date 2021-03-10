// Package main is an example command for Snart.
package main

import (
	"log"
	"os"

	"github.com/go-snart/snart"
	"github.com/go-snart/snart/admin"
)

// tokenEnv is the env var for the bot token.
const tokenEnv = "SNART_TOKEN"

func main() {
	log.SetFlags(log.Flags() | log.Llongfile)

	token, ok := os.LookupEnv(tokenEnv)
	if !ok {
		log.Fatalf("please provide %s", tokenEnv)
	}

	b, err := snart.New(token)
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
