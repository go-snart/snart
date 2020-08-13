package route

import "github.com/go-snart/snart/logs"

// Splitter is the *re2.Regexp used in Split.
var Splitter = MustMatch(`((\x60{1,3})(.*)\2)|(\S+)`)

// Split splits a string using a backtick quoting method.
func Split(s string) []string {
	subj := []rune(s)
	args := []string(nil)

	for {
		m, err := Splitter.FindRunesMatch(subj)
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
