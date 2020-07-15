package token

import "fmt"

func ScanToken() (string, error) {
	_f := "ScanToken"

	fmt.Print("enter your discord token: ")

	tok := ""

	n, err := fmt.Scanln(&tok)
	if err != nil {
		Log.Fatal(_f, err)
	}

	if n == 0 {
		return "", fmt.Errorf("token len %d invalid", n)
	}

	return tok, nil
}
