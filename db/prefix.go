package db

import (
	"errors"
	"fmt"
	"strings"

	dg "github.com/bwmarrin/discordgo"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

var PrefixFail = errors.New("failed to get a prefix")

type Prefix struct {
	Guild string
	Value string
}

var PrefixTable = r.DB("config").TableCreate(
	"prefix",
	r.TableCreateOpts{
		PrimaryKey: "guild",
	},
)

func (d *DB) Prefix(ses *dg.Session, guild, cont string) (string, string, error) {
	_f := "(*DB).Prefix"
	Log.Debugf(_f, "prefix %s", guild)

	d.Once(ConfigDB)
	d.Once(PrefixTable)

	pfxs := make([]Prefix, 0)
	gu := r.Row.Field("guild")
	q := r.DB("config").Table("prefix").Filter(gu.Eq("").Or(gu.Eq(guild)))
	err := q.ReadAll(&pfxs, d)
	if err != nil {
		err = fmt.Errorf("readall &pfxs: %w", err)
		Log.Error(_f, err)
		return "", "", err
	}

	for _, pfx := range pfxs {
		if pfx.Guild == guild && strings.HasPrefix(cont, pfx.Value) {
			return pfx.Value, pfx.Value, nil
		}
	}

	for _, pfx := range pfxs {
		if pfx.Guild == "" && strings.HasPrefix(cont, pfx.Value) {
			return pfx.Value, pfx.Value, nil
		}
	}

	ument := ses.State.User.Mention() + " "
	if strings.HasPrefix(cont, ument) {
		return ument, "@" + ses.State.User.Username + " ", nil
	}

	mment := ""
	g, err := ses.Guild(guild)
	if err != nil {
		err = fmt.Errorf("guild %#v: %w", guild, err)
		Log.Error(_f, err)
		return "", "", err
	}
	for _, member := range g.Members {
		if member.User.ID == ses.State.User.ID {
			mment = member.Mention() + " "

			name := member.Nick
			if name == "" {
				name = member.User.Username
			}

			if strings.HasPrefix(cont, mment) {
				return mment, "@" + name + " ", nil
			}
		}
	}

	return "", "", PrefixFail
}
