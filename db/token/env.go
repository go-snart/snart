package token

import (
	"os"
	"strconv"
)

// EnvName returns the name of the environment variable used to load a token.
func EnvName(idx int) string {
	s := "SNART_TOKEN"

	if idx >= 0 {
		s += "_" + strconv.Itoa(idx)
	}

	return s
}

// EnvTokens returns the tokens listed in the EnvName environment variable.
func EnvTokens() []string {
	tokens := []string(nil)

	for i := -1; ; i++ {
		iName := EnvName(i)

		token, ok := os.LookupEnv(iName)
		if !ok {
			return tokens
		}

		tokens = append(tokens, token)
	}
}
