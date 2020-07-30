package bot_test

import (
	"context"
	"testing"

	"github.com/go-snart/snart/bot"
)

func TestNewBot(t *testing.T) {
	_, err := bot.New(context.Background())
	if err != nil {
		t.Fatal(err)
	}
}
