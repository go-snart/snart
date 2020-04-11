package bot

import r "gopkg.in/rethinkdb/rethinkdb-go.v6"

var ConfigDB = r.DBCreate("config")
