package db_test

import "testing"

func TestDBStart(t *testing.T) {
	_, err := dbDummyStart()
	if err != nil {
		t.Fatal(err)
	}
}

func TestDBStartBad(t *testing.T) {
	db := dbDummy()
	db.Host = "NOt A hOsT dot COM"

	err := db.Start()
	if err == nil {
		t.Fatal("err is nil")
	}
}
