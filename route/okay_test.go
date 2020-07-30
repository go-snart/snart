package route_test

import (
	"testing"

	"github.com/go-snart/snart/route"
)

func TestOkayAnyTrue(t *testing.T) {
	_, _, _, _, _,
		c := ctxDummy("owo")

	ok := route.Any(
		route.False,
		route.False,
		route.True,
		route.False,
	)
	if !ok(c) {
		t.Fatal("!ok(c)")
	}
}

func TestOkayAnyFalse(t *testing.T) {
	_, _, _, _, _,
		c := ctxDummy("owo")

	ok := route.Any(
		route.False,
		route.False,
		route.False,
		route.False,
	)
	if ok(c) {
		t.Fatal("ok(c)")
	}
}

func TestOkayAllTrue(t *testing.T) {
	_, _, _, _, _,
		c := ctxDummy("owo")

	ok := route.All(
		route.True,
		route.True,
		route.True,
		route.True,
	)
	if !ok(c) {
		t.Fatal("!ok(c)")
	}
}

func TestOkayAllFalse(t *testing.T) {
	_, _, _, _, _,
		c := ctxDummy("owo")

	ok := route.All(
		route.True,
		route.True,
		route.False,
		route.True,
	)
	if ok(c) {
		t.Fatal("ok(c)")
	}
}

func TestOkayGuildAdmin(t *testing.T) {
	_, _, _, _, _,
		c := ctxDummy("owo")
	c.Message.ID = "707316494132052008"
	c.Message.ChannelID = "614088693216706590"
	c.Message.GuildID = "466269100214321154"
	c.Message.Author.ID = "304000458144481280"

	ok := route.GuildAdmin(c)
	if !ok {
		t.Fatal("should be ok")
	}
}

func TestOkayGuildAdminBad(t *testing.T) {
	_, _, _, _, _,
		c := ctxDummy("owo")

	ok := route.GuildAdmin(c)
	if ok {
		t.Fatal("shouldn't be ok")
	}
}
