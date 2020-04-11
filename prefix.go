package snart

import (
	"errors"
	"fmt"
	"strings"

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

func (b *Bot) Prefix(guild, cont string) (string, string, error) {
	_f := "(*Bot).GetPrefix"
	Log.Debugf(_f, "prefix %s", guild)

	b.DB.Easy(ConfigDB)
	b.DB.Easy(PrefixTable)

	pfxs := make([]Prefix, 0)
	gu := r.Row.Field("guild")
	q := r.DB("config").Table("prefix").Filter(gu.Eq("").Or(gu.Eq(guild)))
	err := q.ReadAll(&pfxs, b.DB)
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

	ument := b.Session.State.User.Mention() + " "
	if strings.HasPrefix(cont, ument) {
		return ument, "@" + b.Session.State.User.Username + " ", nil
	}

	mment := ""
	g, err := b.Session.Guild(guild)
	if err != nil {
		err = fmt.Errorf("guild %#v: %w", guild, err)
		Log.Error(_f, err)
		return "", "", err
	}
	for _, member := range g.Members {
		if member.User.ID == b.Session.State.User.ID {
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
