package db

import (
	"fmt"
	"sync"

	"github.com/go-snart/snart/logs"
)

var StdinMu = &sync.Mutex{}

// StdinConnStrings gets conn strings from input on the command line.
func StdinConnStrings(name string) []string {
	StdinMu.Lock()
	defer StdinMu.Unlock()

	connStrings := []string(nil)
	connString := ""

	logs.Info.Printf("getting %q conn strings from stdin\n", name)

	for {
		logs.Info.Println("enter a new conn string, or nothing to finish")
		fmt.Print(" > ")

		_, err := fmt.Scanln(&connString)
		if err != nil {
			err = fmt.Errorf("scanln connString: %w", err)
			logs.Warn.Fatalln(err)
		}

		if connString == "" {
			return connStrings
		}

		connStrings = append(connStrings, connString)
	}
}
