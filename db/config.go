package db

import r "gopkg.in/rethinkdb/rethinkdb-go.v6"

var ConfigDB = r.Branch(
	r.DBList().Contains("config"),
	r.Expr(nil),
	r.DBCreate("config"),
).Do(func() r.Term {
	return r.DB("config")
})
