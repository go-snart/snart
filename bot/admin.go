package bot

import (
	"fmt"

	"github.com/go-snart/snart/db"
)

func (b *Bot) adminCacheOwner() {
	_f := "(*Bot).AdminCacheOwner"

	b.WaitReady()

	b.DB.Cache.Lock()
	admin := b.DB.Cache.Get("admin").(db.Cache)
	b.DB.Cache.Unlock()

	app, err := b.Session.Application("@me")
	if err != nil {
		err = fmt.Errorf("app @me: %w", err)
		Log.Warn(_f, err)

		return
	}

	admin.Lock()
	admin.Set(app.Owner.ID, &db.Admin{ID: app.Owner.ID})
	admin.Unlock()
}
