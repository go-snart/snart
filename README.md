[![latest tag](https://img.shields.io/github/v/tag/go-snart/snart)](https://github.com/go-snart/snart/tags)
[![build status](https://img.shields.io/github/workflow/status/go-snart/snart/build)](https://github.com/go-snart/snart/actions?query=workflow:build)
[![lint status](https://img.shields.io/github/workflow/status/go-snart/snart/lint?label=lint)](https://github.com/go-snart/snart/actions?query=workflow:lint)
[![pkg.go.dev docs](https://img.shields.io/badge/pkg.go.dev-docs-blue.svg)](https://pkg.go.dev/github.com/go-snart/snart)
[![open issues](https://img.shields.io/github/issues/go-snart/snart)](https://github.com/go-snart/snart/issues)

Snart
=====
A Discord bot framework in Go. Currently still a little janky, but feel free to use it if you know what you're doing!

Requirements
------------
- A Redis instance

What does Snart mean?
---------------------
Well, it's the word "trans" backwards. It could stand for "Spaghetti Noodles Are Really Tasty". Or it could not. Who knows?

How does it work?
-----------------
Currently, the process is as such:
 - Get the name to use for the Bot:
   - Generally parsed from `os.Args[0]`.
   - The following entries will be using the name "snart".
 - Create a Bot:
   - Open a DB:
     - Check the `SNART_DB(_#)` environment variable(s) for Redis TCP addresses.
     - Also include sane defaults like `snart_db` and `snart_db.docker`.
     - Create a connection pool, and attempt to ping a connection with it.
   - Create a new Handler.
   - Use all unprivileged intents by default.
   - Create a Gamer Queue with an Uptime Gamer.
 - Run the Bot:
   - Create a Halt channel.
   - Use the `admin` and `help` Plugs by default.
   - Load Plugs from Go plugins:
     - Check the `SNART_PLUGS(_#)` environment variable(s) for places to load plugins from.
     - Also include the sane default `plugs` (in current directory).
     - Glob these directories, and use the `plugin` package to load Plugs from files.
   - Iterate all of the Plugs:
     - Plug the DB, Handler, and Halt channel.
     - Get the Intents and Gamers, and add them.
   - Open a Session:
     - Check the `SNART_TOKEN(_#)` environment variable(s) for Discord tokens.
     - Retrieve Discord tokens from the DB (`LRANGE` on the `tokens` key).
     - Try the tokens until a working one is found.
     - Wait for the Discord "ready" event.
   - Add the Handler to the Session.
   - Iterate the Plugs again, and plug the Session.
   - Spawn the Gamer Queue cycling on the Session.
   - Wait for a Halt from the Halt channel, and return it as an error.
