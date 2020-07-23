// Package prefix contains command prefix stuff for db.
package prefix

import (
	"context"
	"errors"
	"fmt"
	"strings"

	dg "github.com/bwmarrin/discordgo"

	"github.com/go-snart/snart/db"
	"github.com/go-snart/snart/db/cache"
)

// ErrPrefixFail is an error which indicates that a function failed to get a prefix.
var ErrPrefixFail = errors.New("failed to get a prefix")

// Table is a table builder for config.admin.
func Table(ctx context.Context, d *db.DB) {
	x, err := d.Conn(&ctx).
		Exec(ctx, `CREATE TABLE IF NOT EXISTS prefix(guild TEXT PRIMARY KEY UNIQUE, value TEXT)`)
	Log.Debugf("Table", "%#v %#v", x, err)
}

// Prefix represents a command prefix Value for a given Guild, as well as a human-readable Clean prefix.
type Prefix struct {
	Guild string `rethinkdb:"guild"`
	Value string `rethinkdb:"value"`
	Clean string `rethinkdb:"-"`
}

// GuildPrefix gets the prefix for a given Guild.
func GuildPrefix(ctx context.Context, d *db.DB, id string) (*Prefix, error) {
	_f := "(*db.DB).GuildPrefix"

	d.Cache.Lock()
	if !d.Cache.Has("prefix") {
		d.Cache.Set("prefix", cache.NewLRUCache(10))
	}

	pfxCache := d.Cache.Get("prefix").(cache.Cache)
	d.Cache.Unlock()

	pfxCache.Lock()

	pfx := pfxCache.Get(id).(*Prefix)
	if pfx != nil {
		return pfx, nil
	}

	pfxCache.Unlock()

	Table(ctx, d)

	const q = `SELECT guild, value FROM prefix WHERE guild == $1`

	rows, err := d.Conn(&ctx).Query(ctx, q, id)
	if err != nil {
		err = fmt.Errorf("db query %#q: %w", q, err)
		Log.Error(_f, err)

		return nil, err
	}

	if rows.Next() {
		pfx = &Prefix{}

		err = rows.Scan(&pfx.Guild, &pfx.Value)
		if err != nil {
			err = fmt.Errorf("scan into pfx: %w", err)
			Log.Error(_f, err)

			return nil, err
		}

		pfxCache.Lock()
		pfxCache.Set(pfx.Guild, pfx)
		pfxCache.Unlock()

		pfx.Clean = pfx.Value

		return pfx, nil
	}

	return nil, ErrPrefixFail
}

// DefaultPrefix gets the default prefix (aka the Guild "").
func DefaultPrefix(ctx context.Context, d *db.DB) (*Prefix, error) {
	return GuildPrefix(ctx, d, "")
}

func userPrefix(ses *dg.Session, cont string, gpfx, dpfx *Prefix) *Prefix {
	ument := ses.State.User.Mention()
	if strings.HasPrefix(cont, ument) {
		pfx := &Prefix{
			Guild: "",
			Value: ument,
		}

		switch {
		case gpfx != nil:
			pfx.Clean = gpfx.Value
		case dpfx != nil:
			pfx.Clean = dpfx.Value
		default:
			pfx.Clean = "@" + ses.State.User.Username + " "
		}

		return pfx
	}

	return nil
}

func memberPrefix(ses *dg.Session, guild, cont string, gpfx, dpfx *Prefix) (*Prefix, error) {
	_f := "memberPrefix"

	mme, err := ses.GuildMember(guild, ses.State.User.ID)
	if err != nil {
		err = fmt.Errorf("member %#v @me: %w", guild, err)
		Log.Error(_f, err)

		return nil, err
	}

	mment := mme.Mention()
	if strings.HasPrefix(cont, mment) {
		pfx := &Prefix{
			Guild: "",
			Value: mment,
		}

		switch {
		case gpfx != nil:
			pfx.Clean = gpfx.Value
		case dpfx != nil:
			pfx.Clean = dpfx.Value
		case mme.Nick != "":
			pfx.Clean = "@" + mme.Nick + " "
		default:
			pfx.Clean = "@" + mme.User.Username + " "
		}

		return pfx, nil
	}

	return nil, nil
}

// FindPrefix finds a matching prefix for a given guild and message content.
func FindPrefix(ctx context.Context, d *db.DB, ses *dg.Session, guild, cont string) (*Prefix, error) {
	_f := "(*db.DB).FindPrefix"

	Log.Debugf(_f, "prefix %s", guild)

	gpfx, err := GuildPrefix(ctx, d, guild)
	if err != nil {
		return nil, err
	}

	if gpfx != nil {
		if strings.HasPrefix(cont, gpfx.Value) {
			return gpfx, nil
		}
	}

	dpfx, err := DefaultPrefix(ctx, d)
	if err != nil {
		return nil, err
	}

	if dpfx != nil {
		if strings.HasPrefix(cont, dpfx.Value) {
			return dpfx, nil
		}
	}

	upfx := userPrefix(ses, cont, gpfx, dpfx)
	if upfx != nil {
		return upfx, nil
	}

	mpfx, err := memberPrefix(ses, guild, cont, gpfx, dpfx)
	if err != nil {
		return nil, err
	}

	if mpfx != nil {
		return mpfx, nil
	}

	return nil, ErrPrefixFail
}
