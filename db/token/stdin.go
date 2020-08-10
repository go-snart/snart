package token

import (
	"fmt"
	"strings"

	"github.com/go-snart/snart/logs"
)

// StdinTokens gets a token from input on the command line.
func StdinTokens() ([]string, error) {
	fmt.Print("enter your discord token(s): ")

	toks := ""

	_, err := fmt.Scanln(&toks)
	if err != nil {
		err = fmt.Errorf("scanln toks: %w", err)

		logs.Warn.Println(err)

		return nil, err
	}

	return strings.Split(toks, ":"), nil
}
