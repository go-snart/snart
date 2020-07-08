package db_test

import "testing"

func TestWaitReady(t *testing.T) {
	db := dbDummy()

	go func() {
		err := db.Start()
		if err != nil {
			panic(err)
		}
	}()

	db.WaitReady()
}
