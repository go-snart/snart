// Package main is an example command for Snart.
package main

import (
	"flag"
	"log"

	"github.com/superloach/confy"

	"github.com/go-snart/snart"

	admin "github.com/go-snart/plug-admin"
)

var (
	confyDir = flag.String("confyDir", ".", "base directory for configuration files")
	confyExt = flag.String("confyExt", ".json", "file extension for configuration files")
)

func main() {
	flag.Parse()

	log.SetFlags(log.Flags() | log.Llongfile)

	c, err := confy.NewOS(*confyDir, *confyExt)
	if err != nil {
		log.Fatalf("confy: %s", err)
	}

	b, err := snart.New(c)
	if err != nil {
		log.Fatalf("open: %s", err)
	}

	err = b.Plug(&admin.Admin{})
	if err != nil {
		log.Fatalf("plug admin: %s", err)
	}

	err = b.Run()
	if err != nil {
		log.Fatalf("run: %s", err)
	}
}
