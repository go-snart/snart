package bot

import (
	"github.com/go-snart/snart/lib/errs"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

type Admin struct {
	UserID string
}

var AdminTable = r.DB("config").TableCreate(
	"admin",
	r.TableCreateOpts{
		PrimaryKey: "userid",
	},
)

func (b *Bot) Admins() ([]string, error) {
	_f := "(*Bot).Admins"

	b.DB.Easy(ConfigDB)
	b.DB.Easy(AdminTable)

	ads := make([]Admin, 0)
	q := r.DB("config").Table("admin")

	err := q.ReadAll(&ads, b.DB)
	if err != nil {
		errs.Wrap(&err, `q.ReadAll(&ads, d)`)
		Log.Error(_f, err)
		return nil, err
	}

	aids := make([]string, len(ads))
	for i, ad := range ads {
		aids[i] = ad.UserID
	}

	return aids, nil
}
