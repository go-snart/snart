// Package main is the command for Snart.
package main

import (
	"flag"

	"github.com/go-snart/snart"
	"github.com/go-snart/snart/logs"
)

var db = flag.String("db", "", "database uri")

func main() {
	flag.Parse()

	if *db == "" {
		logs.Warn.Fatalln("please provide a database uri")
	}

	err := snart.New(*db).Run(ctx)
	if err != nil {
		logs.Warn.Fatalln(err)
	}
}
