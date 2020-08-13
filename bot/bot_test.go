package bot_test

import (
	"testing"

	"github.com/go-snart/snart/test"
)

func TestBot(t *testing.T) {
	bot := test.Bot()
	if bot == nil {
		t.Error("bot == nil")
	}
}
