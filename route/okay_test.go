package route

import (
	"os"
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

	if mid := os.Getenv("SNART_TEST_OKAY_GUILD_ADMIN_MID"); mid != "" {
		c.Message.ID = mid
	}
	if cid := os.Getenv("SNART_TEST_OKAY_GUILD_ADMIN_CID"); cid != "" {
		c.Message.ChannelID = cid
	}
	if gid := os.Getenv("SNART_TEST_OKAY_GUILD_ADMIN_GID"); gid != "" {
		c.Message.GuildID = gid
	}
	if aid := os.Getenv("SNART_TEST_OKAY_GUILD_ADMIN_AID"); aid != "" {
		c.Message.Author.ID = aid
	}

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
