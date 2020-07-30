package route_test

import "testing"

func TestNewFlag(t *testing.T) {
	_, _, _, _, _,
		c := ctxDummy("uwu")
	name, _,
		flag := flagDummy(c)

	if flag.Name() != name {
		t.Fatal("flag.Name() != name")
	}
}

func TestFlagUsage(t *testing.T) {
	_, _, _, _, _,
		c := ctxDummy("uwu")
	_, _,
		flag := flagDummy(c)

	usage := flag.Usage()

	if usage.Content != "" {
		t.Fatal("usage.Content != \"\"")
	}

	if usage.Embed.Title != "Usage of `"+c.Route.Name+"`" {
		t.Fatal("usage.Embed.Title")
	}
}

func TestFlagsParse(t *testing.T) {
	_, _, _, _, _,
		c := ctxDummy("uwu")
	_, args, flag := flagDummy(c)

	err := flag.Parse()
	if err != nil {
		t.Fatal(err)
	}

	fargs := flag.Args()

	if len(args) != len(fargs) {
		t.Fatalf(
			"len(args) == %d != len(fargs) == %d",
			len(args), len(fargs),
		)
	}

	for i := 0; i < len(args) && i < len(fargs); i++ {
		if args[i] != fargs[i] {
			t.Fatalf(
				"args[%d] == %q != fargs[%d] == %q",
				i, args[i],
				i, fargs[i],
			)
		}
	}
}

func TestFlagsParseErr(t *testing.T) {
	_, _, _, _, _,
		c := ctxDummy("uwu")
	_, args, flag := flagDummy(c)
	args[0] = "-badflag"

	err := flag.Parse()
	if err == nil {
		t.Fatal("err is nil")
	}
}

func TestFlagsFlags(t *testing.T) {
	_, _, _, _, _,
		c := ctxDummy("uwu")
	_, args, flag := flagDummy(c)
	args[0] = "-foo"
	args[1] = "bar"
	args[2] = "-baz"

	foo := flag.String("foo", "", "foo flag")
	baz := flag.Bool("baz", false, "baz flag")

	err := flag.Parse()
	if err != nil {
		t.Fatal(err)
	}

	if *foo != "bar" {
		t.Fatalf("*foo == %q != %q", *foo, "bar")
	}

	if !*baz {
		t.Fatal("!*baz")
	}
}

func TestFlagsOutputBuilder(t *testing.T) {
	_, _, _, _, _,
		c := ctxDummy("uwu")
	_, _, flag := flagDummy(c)
	oout := "hello world"

	_, err := flag.FlagSet.Output().Write([]byte(oout))
	if err != nil {
		t.Fatal(err)
	}

	out := flag.Output()
	if out != oout {
		t.Fatalf(
			"out == %q != oout == %q",
			out, oout,
		)
	}
}

type dummyWriter struct{}

func (d dummyWriter) Write(p []byte) (int, error) {
	return 0, nil
}

func TestFlagsOutputOther(t *testing.T) {
	_, _, _, _, _,
		c := ctxDummy("uwu")
	_, _, flag := flagDummy(c)
	oout := "hello world"

	flag.SetOutput(dummyWriter{})

	_, err := flag.FlagSet.Output().Write([]byte(oout))
	if err != nil {
		t.Fatal(err)
	}

	out := flag.Output()
	if len(out) > 0 {
		t.Fatalf("len(out) == %d > 0", len(out))
	}
}
