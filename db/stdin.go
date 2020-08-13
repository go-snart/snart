package db

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"github.com/go-snart/snart/logs"
)

var StdinMu = &sync.Mutex{}

// StdinConnStrings gets conn strings from input on the command line.
func StdinConnStrings(name string) []string {
	StdinMu.Lock()
	defer StdinMu.Unlock()

	connStrings := []string(nil)
	scanner := bufio.NewScanner(os.Stdin)

	logs.Info.Printf("getting %q conn strings from stdin\n", name)

	for {
		logs.Info.Printf("enter a new %q conn string, or nothing to finish\n", name)
		fmt.Print(" > ")

		if !scanner.Scan() {
			break
		}

		connString := scanner.Text()
		if connString == "" {
			return connStrings
		}

		connStrings = append(connStrings, connString)
	}

	err := scanner.Err()
	if err != nil {
		err = fmt.Errorf("scanner err: %w", err)
		logs.Warn.Println(err)
		return nil
	}

	return connStrings
}
