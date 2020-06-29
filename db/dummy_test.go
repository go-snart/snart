package db

import (
	"os"
	"strconv"

	dg "github.com/bwmarrin/discordgo"
)

func dbDummy() *DB {
	db := &DB{
		User: "admin",
		Pass: "",
		Host: "localhost",
		Port: 28015,
	}

	if u := os.Getenv("SNART_DBUSER"); u != "" {
		db.User = u
	}
	if p := os.Getenv("SNART_DBPASS"); p != "" {
		db.Pass = p
	}
	if h := os.Getenv("SNART_DBHOST"); h != "" {
		db.Host = h
	}
	if t, err := strconv.Atoi(os.Getenv("SNART_DBPORT")); err != nil && t > 0 {
		db.Port = t
	}

	return db
}

func dbDummyStart() (*DB, error) {
	db := dbDummy()
	return db, db.Start()
}

func sessionDummy() (
	string,
	*dg.Session,
) {
	tok := os.Getenv("SNART_TOKEN")
	if tok == "" {
		panic("please provide SNART_TOKEN")
	}

	session, err := dg.New(tok)
	if err != nil {
		panic(err)
	}

	err = session.Open()
	if err != nil {
		panic(err)
	}

	return tok, session
}
