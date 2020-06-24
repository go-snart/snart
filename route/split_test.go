package route

import "testing"

func TestSplitErr(t *testing.T) {
	_ = Split("	` ` `	")
}
