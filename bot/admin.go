package bot

import "github.com/go-snart/snart/route"

func (b *Bot) Admin(c *route.Ctx) bool {
	_f := "(*Bot).Admin"

	if route.BotOwner(c) {
		return true
	}

	adminIDs, err := b.DB.AdminIDs()
	if err != nil {
		Log.Warn(_f, err)
		return false
	}

	for _, adminID := range adminIDs {
		if adminID == c.Message.Author.ID {
			return true
		}
	}

	return false
}
