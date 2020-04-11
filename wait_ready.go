package snart

import "time"

func (b *Bot) WaitReady() {
	_f := "(*Bot).WaitReady"

	for {
		Log.Debug(_f, "wait for session")
		if b.Session.State.User != nil {
			break
		}
		time.Sleep(time.Millisecond * 100)
	}

	for {
		Log.Debug(_f, "wait for db")
		if b.DB.Session != nil {
			break
		}
		time.Sleep(time.Millisecond * 100)
	}
}
