package db

import (
	"context"
	"errors"
	"fmt"
	"strings"

	dg "github.com/bwmarrin/discordgo"
	"github.com/gomodule/redigo/redis"

	"github.com/go-snart/snart/log"
)

// Prefix represents a command prefix Value for a given Guild, as well as a human-readable Clean prefix.
type Prefix struct {
	Value string
	Clean string
}

// GuildPrefix gets the prefix for a given guild.
func (d *DB) GuildPrefix(ctx context.Context, id string) (*Prefix, error) {
	conn, err := d.GetContext(ctx)
	if err != nil {
		err = fmt.Errorf("get conn: %w", err)
		log.Warn.Println(err)

		return nil, err
	}
	defer conn.Close()

	pfx, err := redis.String(conn.Do("HGET", id, "prefix"))
	if err != nil {
		if errors.Is(err, redis.ErrNil) {
			return nil, nil
		}

		err = fmt.Errorf("hget %q prefix: %w", id, err)
		log.Warn.Println(err)

		return nil, err
	}

	return &Prefix{
		Value: pfx,
		Clean: pfx,
	}, nil
}

// DefaultPrefix gets the default prefix (aka the guild "").
func (d *DB) DefaultPrefix(ctx context.Context) (*Prefix, error) {
	const id = ""

	pfx, err := d.GuildPrefix(ctx, id)
	if err != nil {
		err = fmt.Errorf("guild prefix %q: %w", id, err)
		log.Warn.Println(err)

		return nil, err
	}

	return pfx, nil
}

func (d *DB) userPrefix(
	ses *dg.Session,
	cont string,
	gpfx, dpfx *Prefix,
) *Prefix {
	ument := ses.State.User.Mention()
	log.Debug.Printf("%q %q", cont, ument)

	if strings.HasPrefix(cont, ument) {
		pfx := &Prefix{
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

func (d *DB) memberPrefix(
	ses *dg.Session,
	guild, cont string,
	gpfx, dpfx *Prefix,
) (*Prefix, error) {
	me := ses.State.User.ID

	mme, err := ses.GuildMember(guild, me)
	if err != nil {
		err = fmt.Errorf("member %q %q (@me): %w", guild, me, err)
		log.Warn.Println(err)

		return nil, err
	}

	mment := mme.Mention()
	if strings.HasPrefix(cont, mment) {
		pfx := &Prefix{
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

// FindPrefix finds a matching prefix for a given guild id and message content.
func (d *DB) FindPrefix(
	ctx context.Context,
	ses *dg.Session,
	guild, cont string,
) (*Prefix, error) {
	log.Debug.Printf("prefix %s", guild)

	gpfx, err := d.GuildPrefix(ctx, guild)
	if err != nil {
		err = fmt.Errorf("guild prefix %q: %w", guild, err)
		log.Warn.Println(err)

		return nil, err
	}

	if gpfx != nil {
		if strings.HasPrefix(cont, gpfx.Value) {
			return gpfx, nil
		}
	}

	dpfx, err := d.DefaultPrefix(ctx)
	if err != nil {
		err = fmt.Errorf("default prefix: %w", err)
		log.Warn.Println(err)

		return nil, err
	}

	if dpfx != nil {
		if strings.HasPrefix(cont, dpfx.Value) {
			return dpfx, nil
		}
	}

	upfx := d.userPrefix(ses, cont, gpfx, dpfx)
	if upfx != nil {
		return upfx, nil
	}

	mpfx, err := d.memberPrefix(ses, guild, cont, gpfx, dpfx)
	if err != nil {
		err = fmt.Errorf("member prefix: %w", err)
		log.Warn.Println(err)

		return nil, err
	}

	if mpfx != nil {
		return mpfx, nil
	}

	return nil, nil
}
