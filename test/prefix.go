package test

import (
	"github.com/go-snart/snart/db/prefix"
	"github.com/go-snart/snart/logs"
)

const (
	PrefixValue = "./"
	PrefixClean = "./"
)

func Prefix() *prefix.Prefix {
	logs.Info.Println("enter")
	defer logs.Info.Println("exit")

	return &prefix.Prefix{
		Value: PrefixValue,
		Clean: PrefixClean,
	}
}
