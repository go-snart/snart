package db

// ConnStrings returns a list of useable redis conn strings.
func ConnStrings(name string) []string {
	return append(
		append(
			EnvConnStrings(name),
			StdinStrings(name+" db")...,
		),
		NamedConnStrings(name)...,
	)
}
