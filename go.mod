module github.com/go-snart/snart

go 1.14

require (
	github.com/go-snart/bot v0.0.0-00010101000000-000000000000
	github.com/go-snart/db v0.0.0-00010101000000-000000000000
	github.com/go-snart/plugin-help v0.0.0-00010101000000-000000000000
	github.com/go-snart/plugin-admin v0.0.0-00010101000000-000000000000
	github.com/namsral/flag v1.7.4-pre
	github.com/superloach/minori v0.0.0-20200401022729-31f6f02808bc
)

replace (
	github.com/go-snart/bot => ../bot
	github.com/go-snart/db => ../db
	github.com/go-snart/plugin-help => ../plugin-help
	github.com/go-snart/plugin-admin => ../plugin-admin
	github.com/go-snart/route => ../route
)
