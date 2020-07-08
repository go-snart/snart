package db_test

import (
	"os"
	"strconv"

	dg "github.com/bwmarrin/discordgo"
	"github.com/go-snart/snart/db"
)

func dbDummy() *db.DB {
	d := &db.DB{
		User: "admin",
		Pass: "",
		Host: "localhost",
		Port: 28015,
	}

	if u := os.Getenv("SNART_DBUSER"); u != "" {
		d.User = u
	}

	if p := os.Getenv("SNART_DBPASS"); p != "" {
		d.Pass = p
	}

	if h := os.Getenv("SNART_DBHOST"); h != "" {
		d.Host = h
	}

	if t, err := strconv.Atoi(os.Getenv("SNART_DBPORT")); err != nil && t > 0 {
		d.Port = t
	}

	return d
}

func dbDummyStart() (*db.DB, error) {
	d := dbDummy()
	return d, d.Start()
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
