package test

import (
	"github.com/go-snart/snart/db"
	"github.com/go-snart/snart/logs"
)

const DBName = "test"

func DB() *db.DB {
	logs.Info.Println("enter")
	defer logs.Info.Println("exit")

	return db.New(DBName)
}
