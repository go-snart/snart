package route

import (
	"strings"
	"testing"
)

func routerDummy() (
	*Route,
	*Router,
) {
	router := NewRouter()

	_, _, _, _, _, _,
		route := routeDummy()
	router.Add(
		route,
	)

	return route,
		router
}

func TestNewRouter(t *testing.T) {
	route, router := routerDummy()

	if router == nil {
		t.Fatal("router == nil")
	}

	if len(*router) != 1 {
		t.Fatal("len(*router) != 1")
	}

	if (*router)[0] != route {
		t.Fatal("(*router)[0] != route")
	}
}

func TestRouterCtx(t *testing.T) {
	_, router := routerDummy()

	var (
		pfx  = "./"
		cpfx = "./"
	)

	_,
		ses := sessionDummy()
	_, _, _, _,
		msg := messageDummy("./yeet bar")
	line := strings.Split(msg.Content, "\n")[0]

	c := router.Ctx(pfx, cpfx, ses, msg, line)
	if c == nil {
		t.Fatal("c == nil")
	}
}

func TestRouterCtxBadMatch(t *testing.T) {
	_, router := routerDummy()
	(*router)[0].Match = "["

	var (
		pfx  = "./"
		cpfx = "./"
	)

	_,
		ses := sessionDummy()
	_, _, _, _,
		msg := messageDummy("owo")
	line := strings.Split(msg.Content, "\n")[0]

	c := router.Ctx(pfx, cpfx, ses, msg, line)
	if c != nil {
		t.Fatal("c != nil")
	}
}

func TestRouterCtxNilOkay(t *testing.T) {
	_, router := routerDummy()
	(*router)[0].Okay = nil

	var (
		pfx  = "./"
		cpfx = "./"
	)

	_,
		ses := sessionDummy()
	_, _, _, _,
		msg := messageDummy("owo")
	line := strings.Split(msg.Content, "\n")[0]

	c := router.Ctx(pfx, cpfx, ses, msg, line)
	if c != nil {
		t.Fatal("c != nil")
	}
}

func TestRouterCtxNoArgs(t *testing.T) {
	_, router := routerDummy()

	var (
		pfx  = "./"
		cpfx = "./"
	)

	_,
		ses := sessionDummy()
	_, _, _, _,
		msg := messageDummy("")
	line := strings.Split(msg.Content, "\n")[0]

	c := router.Ctx(pfx, cpfx, ses, msg, line)
	if c != nil {
		t.Fatal("c != nil")
	}
}
