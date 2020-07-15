package bot

import (
	"context"
	"fmt"

	"github.com/go-snart/snart/db/admin"
	"github.com/go-snart/snart/db/cache"
)

func (b *Bot) adminCache(ctx context.Context) {
	_f := "(*Bot).adminCache"

	b.WaitReady()

	admin.AdminCache(ctx, b.DB)

	b.DB.Cache.Lock()
	adminCache := b.DB.Cache.Get("admin").(cache.Cache)
	b.DB.Cache.Unlock()

	app, err := b.Session.Application("@me")
	if err != nil {
		err = fmt.Errorf("app @me: %w", err)
		Log.Warn(_f, err)

		return
	}

	adminCache.Lock()
	adminCache.Set(app.Owner.ID, true)
	adminCache.Unlock()
}
