// +build !snart_debug

package logs

import "github.com/superloach/nilog"

func debug(_ string) *nilog.Logger {
	return nil
}
