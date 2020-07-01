package db

import (
	"errors"
	"fmt"
	"strings"

	dg "github.com/bwmarrin/discordgo"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

// ErrPrefixFail is an error which indicates that a function failed to get a prefix.
var ErrPrefixFail = errors.New("failed to get a prefix")

// Prefix represents a command prefix Value for a given Guild, as well as a human-readable Clean prefix.
type Prefix struct {
	Guild string `rethinkdb:"guild"`
	Value string `rethinkdb:"value"`
	Clean string `rethinkdb:"-"`
}

// PrefixTable is a table builder for config.prefix.
var PrefixTable = BuildTable(
	ConfigDB, "prefix",
	&r.TableCreateOpts{
		PrimaryKey: "guild",
	}, nil,
)

// GuildPrefix gets the prefix for a given Guild.
func (d *DB) GuildPrefix(id string) (*Prefix, error) {
	_f := "(*DB).GuildPrefix"

	if !d.Cache.Has("prefix") {
		d.Cache.Set("prefix", NewLRUCache(10))
	}

	pfx := d.Cache.Get("prefix").(Cache).Get(id)
	if pfx != nil {
		return pfx.(*Prefix), nil
	}

	pfxs := []*Prefix{}
	q := PrefixTable.Get(id)

	err := q.ReadAll(&pfxs, d)
	if err != nil {
		err = fmt.Errorf("readall &pfxs: %w", err)
		Log.Error(_f, err)
		return nil, err
	}

	if len(pfxs) == 0 {
		return nil, ErrPrefixFail
	}

	d.Cache.Get("prefix").(Cache).Set(id, pfxs[0])
	return pfxs[0], nil
}

// DefaultPrefix gets the default prefix (aka the Guild "").
func (d *DB) DefaultPrefix() (*Prefix, error) {
	return d.GuildPrefix("")
}

// FindPrefix finds a matching prefix for a given guild and message content.
func (d *DB) FindPrefix(ses *dg.Session, guild, cont string) (*Prefix, error) {
	_f := "(*DB).FindPrefix"
	Log.Debugf(_f, "prefix %s", guild)

	gpfx, err := d.GuildPrefix(guild)
	if err != nil {
		return nil, err
	}
	if gpfx != nil {
		return gpfx, nil
	}

	dpfx, err := d.DefaultPrefix()
	if err != nil {
		return nil, err
	}
	if dpfx != nil {
		return dpfx, nil
	}

	ument := ses.State.User.Mention()
	if strings.HasPrefix(cont, ument) {
		return &Prefix{
			Guild: "",
			Value: ument,
			Clean: "@" + ses.State.User.Username,
		}, nil
	}

	mme, err := ses.GuildMember(guild, "@me")
	if err != nil {
		err = fmt.Errorf("member %#v @me: %w", guild, err)
		Log.Error(_f, err)
		return nil, err
	}

	mment := mme.Mention()

	if strings.HasPrefix(cont, mment) {
		if mme.Nick == "" {
			mme.Nick = mme.User.Username
		}

		return &Prefix{
			Guild: guild,
			Value: mment,
			Clean: "@" + mme.Nick,
		}, nil
	}

	return nil, ErrPrefixFail
}
