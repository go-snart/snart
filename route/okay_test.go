package route_test

import (
	"context"
	"testing"

	"github.com/go-snart/snart/route"
	"github.com/go-snart/snart/test"
)

func TestOkay(t *testing.T) {
	ctx := context.Background()
	c := test.Ctx(ctx, "./route")

	t.Run("any-true", func(t *testing.T) {
		if !route.Any(
			route.False,
			route.False,
			route.True,
			route.False,
		)(c) {
			t.Error("!ok")
		}
	})

	t.Run("any-false", func(t *testing.T) {
		if route.Any(
			route.False,
			route.False,
			route.False,
			route.False,
		)(c) {
			t.Error("shouldn't be ok")
		}
	})

	t.Run("all-true", func(t *testing.T) {
		if !route.All(
			route.True,
			route.True,
			route.True,
			route.True,
		)(c) {
			t.Error("should be ok")
		}
	})

	t.Run("all-false", func(t *testing.T) {
		if route.All(
			route.True,
			route.True,
			route.False,
			route.True,
		)(c) {
			t.Error("shouldn't be ok")
		}
	})

	t.Run("guildadmin-true", func(t *testing.T) {
		ctx := context.Background()
		c := test.Ctx(ctx, "./yeet")
		c.Message.ID = "707316494132052008"
		c.Message.ChannelID = "614088693216706590"
		c.Message.GuildID = "466269100214321154"
		c.Message.Author.ID = "304000458144481280"

		if !route.GuildAdmin(c) {
			t.Error("should be ok")
		}
	})

	t.Run("guildadmin-false", func(t *testing.T) {
		ok := route.GuildAdmin(c)
		if ok {
			t.Error("shouldn't be ok")
		}
	})
}
