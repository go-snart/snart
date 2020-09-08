package test

import (
	"context"

	"github.com/go-snart/snart/route"
)

// Flag gets a test *route.Flag.
func Flag(ctx context.Context, content string) *route.Flag {
	return Ctx(ctx, content).Flag
}
