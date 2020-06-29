package route

import (
	"testing"
)

func TestOkayAnyTrue(t *testing.T) {
	_, _, _, _, _, _,
		c := ctxDummy("owo")

	ok := Any(
		False,
		False,
		True,
		False,
	)
	if !ok(c) {
		t.Fatal("!ok(c)")
	}
}

func TestOkayAnyFalse(t *testing.T) {
	_, _, _, _, _, _, c := ctxDummy("owo")

	ok := Any(
		False,
		False,
		False,
		False,
	)
	if ok(c) {
		t.Fatal("ok(c)")
	}
}

func TestOkayAllTrue(t *testing.T) {
	_, _, _, _, _, _, c := ctxDummy("owo")

	ok := All(
		True,
		True,
		True,
		True,
	)
	if !ok(c) {
		t.Fatal("!ok(c)")
	}
}

func TestOkayAllFalse(t *testing.T) {
	_, _, _, _, _, _, c := ctxDummy("owo")

	ok := All(
		True,
		True,
		False,
		True,
	)
	if ok(c) {
		t.Fatal("ok(c)")
	}
}

func TestOkayGuildAdmin(t *testing.T) {
	_, _, _, _, _, _, c := ctxDummy("owo")
	c.Message.ID = "707316494132052008"
	c.Message.ChannelID = "614088693216706590"
	c.Message.GuildID = "466269100214321154"
	c.Message.Author.ID = "304000458144481280"

	ok := GuildAdmin(c)
	if !ok {
		t.Fatal("should be ok")
	}
}

func TestOkayGuildAdminBad(t *testing.T) {
	_, _, _, _, _, _, c := ctxDummy("owo")

	ok := GuildAdmin(c)
	if ok {
		t.Fatal("shouldn't be ok")
	}
}
