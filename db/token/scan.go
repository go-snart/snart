package token

import "fmt"

// ScanToken gets a token from input on the command line.
func ScanToken() (string, error) {
	fmt.Print("enter your discord token: ")

	tok := ""

	_, err := fmt.Scanln(&tok)
	if err != nil {
		return "", err
	}

	return tok, nil
}
