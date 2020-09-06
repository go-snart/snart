package token

import (
	"fmt"

	"github.com/go-snart/snart/db"
	"github.com/go-snart/snart/log"
)

// GetTokens retrieves bot tokens from a DB.
func GetTokens(d *db.DB) []string {
	count, err := d.LLen("tokens").Result()
	if err != nil {
		err = fmt.Errorf("len tokens: %w", err)
		log.Warn.Println(err)

		return nil
	}

	tokens, err := d.LRange("tokens", 0, count).Result()
	if err != nil {
		err = fmt.Errorf("range tokens %d %d: %w", 0, count, err)
		log.Warn.Println(err)

		return nil
	}

	return tokens
}

// StoreTokens adds tokens to the database so that they're persistent.
func StoreTokens(d *db.DB, tokens []string) {
	if len(tokens) == 0 {
		return
	}

	itokens := []interface{}(nil)
	for _, token := range tokens {
		itokens = append(itokens, token)
	}

	_, err := d.LPush("tokens", itokens...).Result()
	if err != nil {
		err = fmt.Errorf("push tokens %v: %w", itokens, err)
		log.Warn.Println(err)
	}
}
