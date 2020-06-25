package db

import "time"

func (d *DB) WaitReady() {
	_f := "(*DB).WaitReady"

	for !d.failed {
		Log.Debug(_f, "wait for db")
		if d.Session != nil {
			break
		}
		time.Sleep(time.Millisecond * 100)
	}
}
