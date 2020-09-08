package route_test

import (
	"context"
	"testing"

	"github.com/go-snart/snart/test"
)

func TestCtx(t *testing.T) {
	ctx := context.Background()
	c := test.Ctx(ctx, "./route a b c")

	t.Run("nil", func(t *testing.T) {
		if c == nil {
			t.Error("c == nil")
		}
	})

	t.Run("run", func(t *testing.T) {
		err := c.Run()
		if err != nil {
			t.Errorf("run c: %w", err)
		}

		if c.Route.Desc != test.RouteDescNew {
			t.Errorf("c.Route.Desc == %q != test.RouteDescNew == %q", c.Route.Desc, test.RouteDescNew)
		}
	})
}
