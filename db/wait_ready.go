package db

import "time"

// WaitReady loops until the DB has a valid Session.
func (d *DB) WaitReady() {
	_f := "(*DB).WaitReady"

	for !d.failed && (d.Session == nil || d.Cache == nil) {
		Log.Debug(_f, "wait for db")

		time.Sleep(time.Second / 10)
	}
}
