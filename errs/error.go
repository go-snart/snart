package errs

import (
	"fmt"
	"os"
)

type Error struct {
	Sig os.Signal
}

func (e Error) Error() string {
	switch {
	case e.Sig != nil:
		return fmt.Sprintf("sig: %s", e.Sig)
	default:
		return "unknown"
	}
}
