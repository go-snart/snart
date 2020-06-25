package bot

import "testing"

func TestNewBot(t *testing.T) {
	_, err := NewBot(nil)
	if err != nil {
		t.Fatal(err)
	}
}
