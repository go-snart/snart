package db

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"github.com/go-snart/snart/logs"
)

var stdinMu = &sync.Mutex{}

func stdinPiped() bool {
	fi, err := os.Stdin.Stat()
	if err != nil {
		err = fmt.Errorf("stat stdin: %w", err)
		logs.Warn.Fatalln(err)

		return false
	}

	return fi.Mode()&os.ModeCharDevice == 0
}

// StdinStrings gets strings from input on the command line.
func StdinStrings(name string) []string {
	if stdinPiped() {
		return nil
	}

	stdinMu.Lock()
	defer stdinMu.Unlock()

	strs := []string(nil)
	scanner := bufio.NewScanner(os.Stdin)

	logs.Info.Printf("getting %q strings from stdin\n", name)

	for {
		logs.Info.Printf("enter a new %q string, or nothing to finish\n", name)
		fmt.Print(" > ")

		if !scanner.Scan() {
			break
		}

		str := scanner.Text()
		if str == "" {
			return strs
		}

		strs = append(strs, str)
	}

	err := scanner.Err()
	if err != nil {
		err = fmt.Errorf("scanner err: %w", err)
		logs.Warn.Println(err)

		return nil
	}

	return strs
}
