package db

// NamedConnStrings returns a selection of simple conn strings for by custom dns (such as in docker compose or stack).
func NamedConnStrings(name string) []string {
	return []string{
		"redis://" + name + "_db.docker",
		"redis://" + name + "_db",
	}
}
