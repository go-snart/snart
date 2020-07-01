package db

import "runtime"

// WaitReady loops until the DB has a valid Session.
func (d *DB) WaitReady() {
	_f := "(*DB).WaitReady"

	for !d.failed {
		Log.Debug(_f, "wait for db")
		if d.Session != nil {
			break
		}

		runtime.Gosched()
	}
}
