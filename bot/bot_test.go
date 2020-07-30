package bot_test

import (
	"testing"

	"github.com/go-snart/snart/bot"
)

func TestNewBot(t *testing.T) {
	b := bot.New()
	if b == nil {
		t.Fatal("nil bot")
	}
}
