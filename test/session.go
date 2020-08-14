package test

import (
	dg "github.com/bwmarrin/discordgo"

	"github.com/go-snart/snart/db/token"
)

// SessionDB is a cache *db.DB for Session.
var SessionDB = DB()

// Session gets a test *dg.Session.
func Session() *dg.Session {
	return token.Open(SessionDB, nil)
}
