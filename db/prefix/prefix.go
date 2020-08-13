// Package prefix contains command prefix stuff for db.
package prefix

import (
	"errors"
	"fmt"
	"strings"

	dg "github.com/bwmarrin/discordgo"

	"github.com/go-snart/snart/db"
	"github.com/go-snart/snart/logs"
)

// ErrPrefixFail is an error which indicates that a function failed to get a prefix.
var ErrPrefixFail = errors.New("failed to get a prefix")

// Prefix represents a command prefix Value for a given Guild, as well as a human-readable Clean prefix.
type Prefix struct {
	Guild string
	Value string
	Clean string
}

// GuildPrefix gets the prefix for a given Guild.
func GuildPrefix(d *db.DB, id string) (*Prefix, error) {
	pfx, err := d.HGet(id, "prefix").Result()
	if err != nil {
		err = fmt.Errorf("hget %q prefix: %w", id, err)
		logs.Warn.Println(err)
		return nil, err
	}

	return &Prefix{
		Guild: id,
		Value: pfx,
		Clean: pfx,
	}
}

// DefaultPrefix gets the default prefix (aka the Guild "").
func DefaultPrefix(d *db.DB) (*Prefix, error) {
	pfx, err := GuildPrefix(d, "")
	if err != nil {
		err = fmt.Errorf("guild prefix %q: %w", "", err)
		logs.Warn.Println(err)
		return nil, err
	}

	return pfx, nil
}

func userPrefix(ses *dg.Session, cont string, gpfx, dpfx *Prefix) *Prefix {
	ument := ses.State.User.Mention()
	logs.Debug.Printf("%q %q", cont, ument)

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
	me := ses.State.User.ID

	mme, err := ses.GuildMember(guild, me)
	if err != nil {
		err = fmt.Errorf("member %q %q (@me): %w", guild, me, err)

		logs.Warn.Println(err)

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
func FindPrefix(d *db.DB, ses *dg.Session, guild, cont string) (*Prefix, error) {
	logs.Debug.Printf("prefix %s", guild)

	gpfx, err := GuildPrefix(d, guild)
	if err != nil && !errors.Is(err, ErrPrefixFail) {
		err = fmt.Errorf("guild prefix %q: %w", guild, err)

		logs.Warn.Println(err)

		return nil, err
	}

	if gpfx != nil {
		if strings.HasPrefix(cont, gpfx.Value) {
			return gpfx, nil
		}
	}

	dpfx, err := DefaultPrefix(d)
	if err != nil && !errors.Is(err, ErrPrefixFail) {
		err = fmt.Errorf("default prefix: %w", err)

		logs.Warn.Println(err)

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

		logs.Warn.Println(err)

		return nil, err
	}

	if mpfx != nil {
		return mpfx, nil
	}

	return nil, ErrPrefixFail
}
