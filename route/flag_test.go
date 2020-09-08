package route_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/go-snart/snart/test"
)

func TestFlag(t *testing.T) {
	ctx := context.Background()
	flag := test.Flag(ctx, "./route a `b c` ``d ` e`` ```f `` g ` h```")

	t.Run("parse", func(t *testing.T) {
		err := flag.Parse()
		if err != nil {
			t.Errorf("flag parse: %w", err)
		}
	})

	t.Run("name", func(t *testing.T) {
		if flag.Name() != "route" {
			t.Errorf(`flag.Name() == %q != "route"`, flag.Name())
		}
	})

	t.Run("args", func(t *testing.T) {
		args := flag.Args()
		eargs := []string{"a", "b c", "d ` e", "f `` g ` h"}

		if !reflect.DeepEqual(args, eargs) {
			t.Errorf("args == %#v != eargs == %#v", args, eargs)
		}
	})

	t.Run("usage", func(t *testing.T) {
		usage := flag.Usage()

		const econtent = ""
		if usage.Content != econtent {
			t.Errorf("usage.Content == %q != econtent == %q", usage.Content, econtent)
		}

		const etitle = "Usage of `route`"
		if usage.Embed.Title != etitle {
			t.Errorf("usage.Embed.Title == %q != etitle == %q", usage.Embed.Title, etitle)
		}
	})
}

func TestFlagsFlags(t *testing.T) {
	ctx := context.Background()
	flag := test.Flag(ctx, "./route -foo bar -baz")

	foo := flag.String("foo", "", "foo flag")
	baz := flag.Bool("baz", false, "baz flag")

	t.Run("parse", func(t *testing.T) {
		err := flag.Parse()
		if err != nil {
			t.Errorf("flag parse: %w", err)
		}
	})

	t.Run("flags", func(t *testing.T) {
		if *foo != "bar" {
			t.Fatalf("*foo == %q != %q", *foo, "bar")
		}

		if !*baz {
			t.Fatal("!*baz")
		}
	})
}
