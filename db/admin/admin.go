package admin

import (
	"context"
	"fmt"

	"github.com/go-snart/snart/route"
	"github.com/go-snart/snart/db"
	"github.com/go-snart/snart/db/cache"
)

var Log = db.Log.GetLogger("admin")

// AdminTable is a table builder for config.admin.
func AdminTable(ctx context.Context, d *db.DB) {
	x, err := d.Exec(ctx, `CREATE TABLE IF NOT EXISTS admin(id TEXT PRIMARY KEY UNIQUE)`)
	Log.Debugf("admintable", "%#v %#v", x, err)
}

// Admin checks if the author has bot-wide admin privileges.
func Admin(d *db.DB) route.Okay {
	return func(c *route.Ctx) bool {
		d.Cache.Lock()
		adminCache := d.Cache.Get("admin").(cache.Cache)
		d.Cache.Unlock()

		adminCache.Lock()
		defer adminCache.Unlock()

		return adminCache.Has(c.Message.Author.ID)
	}
}

// AdminCache retrieves known admins and puts them into cache.
func AdminCache(ctx context.Context, d *db.DB) {
	_f := "(*db.DB).AdminCache"

	AdminTable(ctx, d)

	const q = `SELECT (id) FROM admin`

	rows, err := d.Query(ctx, q)
	if err != nil {
		err = fmt.Errorf("db query %#q: %w", q, err)
		Log.Warn(_f, err)

		return
	}
	defer rows.Close()

	d.Cache.Lock()

	if !d.Cache.Has("admin") {
		d.Cache.Set("admin", cache.NewMapCache())
	}

	adminCache := d.Cache.Get("admin").(cache.Cache)

	d.Cache.Unlock()

	for rows.Next() {
		admin := ""

		err = rows.Scan(&admin)
		if err != nil {
			err = fmt.Errorf("scan admin: %w", err)
			Log.Warn(_f, err)

			return
		}

		adminCache.Lock()
		adminCache.Set(admin, true)
		adminCache.Unlock()
	}

	err = rows.Err()
	if err != nil {
		err = fmt.Errorf("rows err: %w", err)
		Log.Warn(_f, err)

		return
	}
}
