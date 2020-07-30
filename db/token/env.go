package token

import (
	"errors"
	"os"
	"strings"
)

// EnvTokensName is the name of the environment variable used to load tokens.
//nolint:gosec
const EnvTokensName = "SNART_TOKENS"

// ErrEnvTokensUnset occurs when the EnvTokensName environment variable is not set.
var ErrEnvTokensUnset = errors.New(EnvTokensName + " is not set")

// EnvTokens returns the tokens listed in the EnvTokensName environment variable.
func EnvTokens() ([]string, error) {
	toks, ok := os.LookupEnv(EnvTokensName)
	if !ok {
		return nil, ErrEnvTokensUnset
	}

	return strings.Split(toks, ":"), nil
}
