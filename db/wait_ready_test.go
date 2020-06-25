package db

import "testing"

func TestWaitReady(t *testing.T) {
	db := dbDummy()

	go func() {
		err := db.Start()
		if err != nil {
			t.Fatal(err)
		}
	}()

	db.WaitReady()
}
