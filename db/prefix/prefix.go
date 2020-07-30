// Package prefix contains command prefix stuff for db.
package prefix

import (
	"context"
	"errors"
	"fmt"
	"strings"

	dg "github.com/bwmarrin/discordgo"

	"github.com/go-snart/snart/db"
)

// ErrPrefixFail is an error which indicates that a function failed to get a prefix.
var ErrPrefixFail = errors.New("failed to get a prefix")

// Table is a table builder for config.admin.
func Table(ctx context.Context, d *db.DB) {
	const (
		_f = "Table"
		e  = `CREATE TABLE IF NOT EXISTS prefix(
			guild TEXT PRIMARY KEY UNIQUE,
			value TEXT
		)`
	)

	_, err := d.Conn(&ctx).Exec(ctx, e)
	if err != nil {
		err = fmt.Errorf("exec %#q: %w", e, err)

		Log.Error(_f, err)

		return
	}
}

// Prefix represents a command prefix Value for a given Guild, as well as a human-readable Clean prefix.
type Prefix struct {
	Guild string
	Value string
	Clean string
}

// GuildPrefix gets the prefix for a given Guild.
func GuildPrefix(ctx context.Context, d *db.DB, id string) (*Prefix, error) {
	const _f = "GuildPrefix"

	Table(ctx, d)

	const q = `SELECT guild, value FROM prefix WHERE guild = $1`

	rows, err := d.Conn(&ctx).Query(ctx, q, id)
	if err != nil {
		err = fmt.Errorf("query %#q (%q): %w", q, id, err)

		Log.Error(_f, err)

		return nil, err
	}

	if rows.Next() {
		pfx := &Prefix{}

		err = rows.Scan(&pfx.Guild, &pfx.Value)
		if err != nil {
			err = fmt.Errorf("scan into pfx: %w", err)

			Log.Error(_f, err)

			return nil, err
		}

		pfx.Clean = pfx.Value

		return pfx, nil
	}

	return nil, ErrPrefixFail
}

// DefaultPrefix gets the default prefix (aka the Guild "").
func DefaultPrefix(ctx context.Context, d *db.DB) (*Prefix, error) {
	const _f = "DefaultPrefix"

	pfx, err := GuildPrefix(ctx, d, "")
	if err != nil {
		err = fmt.Errorf("guild prefix %q: %w", "", err)

		Log.Error(_f, err)

		return nil, err
	}

	return pfx, nil
}

func userPrefix(ses *dg.Session, cont string, gpfx, dpfx *Prefix) *Prefix {
	const _f = "userPrefix"

	ument := ses.State.User.Mention()
	Log.Debugf(_f, "%q %q", cont, ument)

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
	const _f = "memberPrefix"

	mme, err := ses.GuildMember(guild, ses.State.User.ID)
	if err != nil {
		err = fmt.Errorf("member %q @me: %w", guild, err)

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
	const _f = "(*db.DB).FindPrefix"

	Log.Debugf(_f, "prefix %s", guild)

	gpfx, err := GuildPrefix(ctx, d, guild)
	if err != nil && !errors.Is(err, ErrPrefixFail) {
		err = fmt.Errorf("guild prefix %q: %w", guild, err)

		Log.Error(_f, err)

		return nil, err
	}

	if gpfx != nil {
		if strings.HasPrefix(cont, gpfx.Value) {
			return gpfx, nil
		}
	}

	dpfx, err := DefaultPrefix(ctx, d)
	if err != nil && !errors.Is(err, ErrPrefixFail) {
		err = fmt.Errorf("default prefix: %w", err)

		Log.Error(_f, err)

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
	if err != nil && !errors.Is(err, ErrPrefixFail) {
		err = fmt.Errorf("member prefix: %w", err)

		Log.Error(_f, err)

		return nil, err
	}

	if mpfx != nil {
		return mpfx, nil
	}

	return nil, ErrPrefixFail
}
