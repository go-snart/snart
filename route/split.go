package route

import (
	re2 "github.com/dlclark/regexp2"

	"github.com/go-snart/snart/logs"
)

var splitter = re2.MustCompile(`((\x60+)(.*)\2)|(\S+)`, 0)

// Split splits a string using a backtick quoting method.
func Split(s string) []string {
	subj := []rune(s)
	args := []string{}

	for {
		m, err := splitter.FindRunesMatch(subj)
		if err != nil {
			logs.Warn.Println(err)
			break
		}

		if m == nil {
			break
		}

		gs := m.Groups()

		match := gs[4].Capture.String()
		if match == "" {
			match = gs[3].Capture.String()
		}

		args = append(args, match)

		l := m.Group.Capture.Length + 1
		if l > len(subj) {
			break
		}

		subj = subj[l:]
	}

	return args
}
