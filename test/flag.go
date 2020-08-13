package test

import (
	"github.com/go-snart/snart/logs"
	"github.com/go-snart/snart/route"
)

func Flag(content string) *route.Flag {
	logs.Info.Println("enter")
	defer logs.Info.Println("exit")

	return Ctx(content).Flag
}
