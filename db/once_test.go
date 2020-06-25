package db

import (
	"testing"

	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

func TestOnceTwice(t *testing.T) {
	db, err := dbDummyStart()
	if err != nil {
		t.Fatal(err)
	}

	db.Once(r.Expr(nil))
	db.Once(r.Expr(nil))
}
