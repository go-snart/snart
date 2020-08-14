package route

import re2 "github.com/dlclark/regexp2"

// Splitter is the *re2.Regexp used in Split.
var Splitter = re2.MustCompile(`(\x60+)(.*?)\1|(\S+)`, 0)

// Split splits a string using a backtick quoting method.
func Split(s string) []string {
	subj := []rune(s)
	args := []string(nil)

	for {
		m, err := Splitter.FindRunesMatch(subj)
		if err != nil {
			break
		}

		if m == nil {
			break
		}

		gs := m.Groups()

		match := gs[3].Capture.String()
		if match == "" {
			match = gs[2].Capture.String()
		}

		args = append(args, match)

		l := gs[0].Capture.Length + 1
		if l > len(subj) {
			break
		}

		subj = subj[l:]
	}

	return args
}
