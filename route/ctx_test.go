package route_test

import "testing"

func TestNewCtx(t *testing.T) {
	pfx, session, message, flag, route,
		c := ctxDummy("owo")

	if c.Prefix.Value != pfx.Value {
		t.Fatalf(
			"c.Prefix.Value == %q != pfx.Value == %q",
			c.Prefix.Value, pfx.Value,
		)
	}

	if c.Prefix.Clean != pfx.Clean {
		t.Fatalf(
			"c.Prefix.Clean == %q != pfx.Clean == %q",
			c.Prefix.Clean, pfx.Clean,
		)
	}

	if c.Session != session {
		t.Fatalf(
			"c.Session == %#v != session == %#v",
			c.Session, session,
		)
	}

	if c.Message != message {
		t.Fatalf(
			"c.Message == %#v != message == %#v",
			c.Message, message,
		)
	}

	if c.Flag != flag {
		t.Fatalf(
			"c.Flag == %#v != flag == %#v",
			c.Flag, flag,
		)
	}

	if c.Route != route {
		t.Fatalf(
			"c.Route == %#v != route == %#v",
			c.Route, route,
		)
	}
}

func TestCtxRun(t *testing.T) {
	_, _, _, _, _,
		c := ctxDummy("uwu")

	err := c.Run()
	if err != nil {
		t.Fatal(err)
	}

	if c.Route.Desc != "run" {
		t.Fatalf("c.Route.Desc != run")
	}
}
