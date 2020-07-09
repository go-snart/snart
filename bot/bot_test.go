package bot_test

import (
	"testing"

	"github.com/go-snart/snart/bot"
)

func TestNewBot(t *testing.T) {
	_, err := bot.NewBot(nil)
	if err != nil {
		t.Fatal(err)
	}
}
