package db

import (
	"os"
	"strconv"
	"strings"
)

// EnvName returns the name of the environment variable used to load a conn string.
func EnvName(name string, idx int) string {
	s := strings.ToUpper(name) + "_DB"

	if idx >= 0 {
		s += "_" + strconv.Itoa(idx)
	}

	return s
}

// EnvConnStrings returns the conn strings listed in the EnvName environment variable.
func EnvConnStrings(name string) []string {
	connStrings := []string(nil)

	for i := -1; ; i++ {
		iName := EnvName(name, i)

		connString, ok := os.LookupEnv(iName)
		if !ok {
			return connStrings
		}

		connStrings = append(connStrings, connString)
	}
}
