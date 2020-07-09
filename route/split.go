package route

import re2 "github.com/dlclark/regexp2"

var splitter = re2.MustCompile(`((\x60+)(.+?)\2)|(\S+)`, 0)

// Split splits a string using a backtick quoting method.
func Split(s string) []string {
	subj := []rune(s)
	args := make([]string, 0)

	for {
		m, _ := splitter.FindRunesMatch(subj)
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
