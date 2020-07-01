package db

import r "gopkg.in/rethinkdb/rethinkdb-go.v6"

func BuildDB(name interface{}) r.Term {
	return r.Branch(
		r.DBList().Contains(name),
		r.Expr(nil),
		r.DBCreate(name),
	).Do(func() r.Term {
		return r.DB(name)
	})
}

func BuildTable(db r.Term, name interface{}, co *r.TableCreateOpts, o *r.TableOpts) r.Term {
	cos := []r.TableCreateOpts{}
	if co != nil {
		cos = append(cos, *co)
	}

	os := []r.TableOpts{}
	if o != nil {
		os = append(os, *o)
	}

	return r.Branch(
		db.TableList().Contains(name),
		r.Expr(nil),
		db.TableCreate(name, cos...),
	).Do(func() r.Term {
		return db.Table(name, os...)
	})
}
