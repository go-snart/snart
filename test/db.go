package test

import "github.com/go-snart/snart/db"

// DBName is the name to use for loading DB configs.
const DBName = "test"

// DB gets a test *db.DB.
func DB() *db.DB {
	return db.New(DBName)
}
