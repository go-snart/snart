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

	if u := os.Getenv("SNART_TEST_DB_USER"); u != "" {
		db.User = u
	}
	if p := os.Getenv("SNART_TEST_DB_PASS"); p != "" {
		db.Pass = p
	}
	if h := os.Getenv("SNART_TEST_DB_HOST"); h != "" {
		db.Host = h
	}
	if t, err := strconv.Atoi(os.Getenv("SNART_TEST_DB_PORT")); err != nil && t > 0 {
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
	tok := os.Getenv("SNART_TEST_TOKEN")
	if tok == "" {
		panic("please provide $SNART_TEST_TOKEN")
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
