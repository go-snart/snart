package db

import (
	"testing"

	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

func TestAdminIDs(t *testing.T) {
	db, err := dbDummyStart()
	if err != nil {
		t.Fatal(err)
	}

	wr, err := r.
		DB("config").Table("admin").
		Insert(Admin{ID: "foobar"}).
		RunWrite(db)
	if err != nil {
		t.Fatal(err)
	}
	if wr.Errors > 0 {
		t.Fatalf(
			"%q and %d more",
			wr.FirstError, wr.Errors-1,
		)
	}

	aids, err := db.AdminIDs()
	if err != nil {
		t.Fatal(err)
	}

	ok := false
	for _, aid := range aids {
		if aid == "foobar" {
			ok = true
		}
	}
	if !ok {
		t.Fatal("foobar not in aids")
	}

	_, err = r.
		DB("config").Table("admin").
		Get("foobar").Delete().
		Run(db)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAdminIDsBad(t *testing.T) {
	db, err := dbDummyStart()
	if err != nil {
		t.Fatal(err)
	}
	db.Close()

	_, err = db.AdminIDs()
	if err == nil {
		t.Fatalf("err shouldn't be nil")
	}
}
