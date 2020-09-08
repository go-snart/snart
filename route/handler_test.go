package route_test

import (
	"testing"

	"github.com/go-snart/snart/test"
)

func TestNewHandler(t *testing.T) {
	handler := test.Handler()

	t.Run("nil", func(t *testing.T) {
		if handler == nil {
			t.Error("handler == nil")
		}
	})

	t.Run("len", func(t *testing.T) {
		if len(handler.Routes) != 1 {
			t.Errorf("len(handler.Routes) == %d != 1", len(handler.Routes))
		}
	})

	t.Run("[0].name", func(t *testing.T) {
		name := handler.Routes[0].Name
		const ename = "route"

		if name != ename {
			t.Errorf("name == %q != ename == %q", name, ename)
		}
	})
}

/*
func TestHandlerCtx(t *testing.T) {
	_,
		handler := handlerDummy()

	_, _,
		pfx := prefixDummy()

	ses := sessionDummy

	_, _, _, _,
		msg := messageDummy("./yeet bar")

	line := strings.Split(msg.Content, "\n")[0]

	c := handler.Ctx(pfx, ses, msg, line)
	if c == nil {
		t.Fatal("c == nil")
	}
}

func TestHandlerCtxBadMatch(t *testing.T) {
	_,
		handler := handlerDummy()
	(*handler)[0].Match = "["

	_, _,
		pfx := prefixDummy()

	ses := sessionDummy

	_, _, _, _,
		msg := messageDummy("owo")

	line := strings.Split(msg.Content, "\n")[0]

	c := handler.Ctx(pfx, ses, msg, line)
	if c != nil {
		t.Fatal("c != nil")
	}
}

func TestHandlerCtxNilOkay(t *testing.T) {
	_,
		handler := handlerDummy()
	(*handler)[0].Okay = nil

	_, _,
		pfx := prefixDummy()

	ses := sessionDummy

	_, _, _, _,
		msg := messageDummy("yeet")

	line := strings.Split(msg.Content, "\n")[0]

	c := handler.Ctx(pfx, ses, msg, line)
	if c == nil {
		t.Fatal("c == nil")
	}

	if !c.Route.Okay((*route.Ctx)(nil)) {
		t.Fatal("c.Route.Okay != True")
	}
}

func TestHandlerCtxNoArgs(t *testing.T) {
	_,
		handler := handlerDummy()

	_, _,
		pfx := prefixDummy()

	ses := sessionDummy

	_, _, _, _,
		msg := messageDummy("")

	line := strings.Split(msg.Content, "\n")[0]

	c := handler.Ctx(pfx, ses, msg, line)
	if c != nil {
		t.Fatal("c != nil")
	}
}

func TestHandlerCtxIndex1(t *testing.T) {
	_,
		handler := handlerDummy()

	_, _,
		pfx := prefixDummy()

	ses := sessionDummy

	_, _, _, _,
		msg := messageDummy("ayeet")

	line := strings.Split(msg.Content, "\n")[0]

	c := handler.Ctx(pfx, ses, msg, line)
	if c != nil {
		t.Fatal("c != nil")
	}
}
*/
