package db

import "time"

func (d *DB) WaitReady() {
	_f := "(*DB).WaitReady"

	for {
		Log.Debug(_f, "wait for db")
		if d.Session != nil {
			break
		}
		time.Sleep(time.Millisecond * 100)
	}
}
