package db_test

import (
	"testing"

	"github.com/go-snart/snart/db"
)

func TestAdminIDs(t *testing.T) {
	d, err := dbDummyStart()
	if err != nil {
		t.Fatal(err)
	}

	wr, err := db.AdminTable.
		Insert(db.Admin{ID: "foobar"}).
		RunWrite(d)
	if err != nil {
		t.Fatal(err)
	}

	if wr.Errors > 0 {
		t.Fatalf(
			"%q and %d more",
			wr.FirstError, wr.Errors-1,
		)
	}

	aids, err := d.AdminIDs()
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

	_, err = db.AdminTable.
		Get("foobar").Delete().
		Run(d)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAdminIDsBad(t *testing.T) {
	d, err := dbDummyStart()
	if err != nil {
		t.Fatal(err)
	}

	d.Close()

	_, err = d.AdminIDs()
	if err == nil {
		t.Fatalf("err shouldn't be nil")
	}
}
