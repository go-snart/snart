package route

import (
	"testing"
)

func TestNewCtx(t *testing.T) {
	prefix, cleanPrefix, session, message, flags, route,
		c := ctxDummy("owo")

	if c.Prefix != prefix {
		t.Fatalf(
			"c.Prefix == %q; prefix == %q",
			c.Prefix, prefix,
		)
	}

	if c.CleanPrefix != cleanPrefix {
		t.Fatalf(
			"c.CleanPrefix == %q; cleanPrefix == %q",
			c.CleanPrefix, cleanPrefix,
		)
	}

	if c.Session != session {
		t.Fatalf(
			"c.Session == %v; session == %v",
			c.Session, session,
		)
	}

	if c.Message != message {
		t.Fatalf(
			"c.Message == %v; message == %v",
			c.Message, message,
		)
	}

	if c.Flags != flags {
		t.Fatalf(
			"c.Flags == %v; flags == %v",
			c.Flags, flags,
		)
	}

	if c.Route != route {
		t.Fatalf(
			"c.Route == %v; route == %v",
			c.Route, route,
		)
	}
}

func TestCtxRun(t *testing.T) {
	_, _, _, _, _, _,
		c := ctxDummy("uwu")

	err := c.Run()
	if err != nil {
		t.Fatal(err)
	}

	if c.Route.Desc != "run" {
		t.Fatalf("c.Route.Desc != %q", "run")
	}
}
