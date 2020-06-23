package db

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-snart/snart/route"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

func (d *DB) Queryer(getq func(*route.Ctx) (r.Term, error)) func(*route.Ctx) error {
	_f := "Queryer"

	return func(ctx *route.Ctx) error {
		rep := ctx.Reply()

		q, err := getq(ctx)
		if err != nil {
			err = fmt.Errorf("apply f to ctx: %w", err)

			rep.Content = err.Error()
			rep.Send()

			Log.Error(_f, err)

			return err
		}
		if q.String() == r.Expr(nil).String() {
			return nil
		}

		var tmp []interface{}
		err = q.ReadAll(&tmp, d)
		if err != nil {
			err = fmt.Errorf("readall tmp %s: %w", q, err)

			rep.Content = err.Error()
			rep.Send()

			Log.Error(_f, err)

			return err
		}

		cont, err := json.MarshalIndent(tmp, "", "\t")
		if err != nil {
			err = fmt.Errorf("marshal cont: %w", q, err)

			rep.Content = err.Error()
			rep.Send()

			Log.Error(_f, err)

			return err
		}

		rep.Content = "```json\n"
		for _, line := range strings.Split(string(cont), "\n") {
			rep.Content += line + "\n"
			if len(rep.Content) > 1950 {
				rep.Content += "```"
				err = rep.Send()
				if err != nil {
					return err
				}
				rep.Content = "```json\n"
			}
		}
		rep.Content += "```"
		return rep.Send()
	}
}
