package route_test

import (
	"testing"

	"github.com/go-snart/snart/test"
)

const content = "./route a b c"

func TestCtx(t *testing.T) {
	c := test.Ctx(content)

	if c == nil {
		t.Error("c == nil")
	}
}

func TestCtxRun(t *testing.T) {
	c := test.Ctx(content)

	err := c.Run()
	if err != nil {
		t.Fatal(err)
	}

	if c.Route.Desc != "run" {
		t.Fatalf("c.Route.Desc != run")
	}
}
