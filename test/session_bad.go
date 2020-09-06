package test

import dg "github.com/bwmarrin/discordgo"

// SessionBadToken is a discord token to use for SessionBad.
const SessionBadToken = "session bad token"

// SessionBad gets a bad test *dg.Session.
func SessionBad() *dg.Session {
	session, _ := dg.New(SessionBadToken)

	return session
}
