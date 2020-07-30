package db

import (
	"errors"
	"os"
	"strings"
)

// EnvName is the name of the environment variable used to load configs.
const EnvName = "SNART_CONFIGS"

// ErrEnvUnset occurs when the EnvName environment variable is not set.
var ErrEnvUnset = errors.New(EnvName + " is not set")

// EnvConfigs returns the configs listed in the EnvName environment variable.
func EnvConfigs() ([]string, error) {
	toks, ok := os.LookupEnv(EnvName)
	if !ok {
		return nil, ErrEnvUnset
	}

	return strings.Split(toks, ":"), nil
}
