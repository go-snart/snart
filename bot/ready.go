package bot

import (
	"time"

	"github.com/go-snart/snart/logs"
)

// WaitReady loops until the Bot has received a Ready event.
func (b *Bot) WaitReady() {
	for !b.Ready {
		logs.Debug.Println("wait for ready")

		time.Sleep(time.Second / 10)
	}
}
