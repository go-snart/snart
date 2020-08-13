package token

import (
	"bufio"
	"fmt"
	"os"

	"github.com/go-snart/snart/db"
	"github.com/go-snart/snart/logs"
)

// StdinTokens gets discord tokens from input on the command line.
func StdinTokens() []string {
	db.StdinMu.Lock()
	defer db.StdinMu.Unlock()

	tokens := []string(nil)
	scanner := bufio.NewScanner(os.Stdin)

	logs.Info.Println("getting discord tokens from stdin")

	for {
		logs.Info.Println("enter a new discord token, or nothing to finish")
		fmt.Print(" > ")

		if !scanner.Scan() {
			break
		}

		token := scanner.Text()
		if token == "" {
			return tokens
		}

		tokens = append(tokens, token)
	}

	err := scanner.Err()
	if err != nil {
		err = fmt.Errorf("scanner err: %w", err)
		logs.Warn.Println(err)
		return nil
	}

	return tokens
}
