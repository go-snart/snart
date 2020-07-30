package token

import (
	"fmt"
	"strings"
)

// StdinTokens gets a token from input on the command line.
func StdinTokens() ([]string, error) {
	const _f = "StdinTokens"

	fmt.Print("enter your discord token(s): ")

	toks := ""

	_, err := fmt.Scanln(&toks)
	if err != nil {
		err = fmt.Errorf("scanln toks: %w", err)

		Log.Error(_f, err)

		return nil, err
	}

	return strings.Split(toks, ":"), nil
}
