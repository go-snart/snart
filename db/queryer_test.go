package db

import (
	"testing"

	"github.com/go-snart/snart/route"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

func TestQueryer(t *testing.T) {
	db, err := dbDummyStart()
	if err != nil {
		t.Fatal(err)
	}

	_ = db.Queryer(func(c *route.Ctx) (r.Term, error) {
		return r.Expr("hello"), nil
	})
}
