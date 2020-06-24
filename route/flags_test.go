package route

import (
	"flag"
	"fmt"
	"testing"

	dg "github.com/bwmarrin/discordgo"
)

func TestNewFlags(t *testing.T) {
	_, _, _, _, _, _,
		c := ctxDummy("uwu")
	name, _,
		flags := flagsDummy(c)

	if flags.ctx != c {
		t.Fatal("flags.ctx != c")
	}
	if flags.Name() != name {
		t.Fatal("flags.Name() != name")
	}

}

func TestFlagsUsage(t *testing.T) {
	_, _, _, _, _, _,
		c := ctxDummy("uwu")
	_, _,
		flags := flagsDummy(c)

	usage := flags.Usage()

	if usage.Content != "" {
		t.Fatal("usage.Content != \"\"")
	}
	if usage.Embed.Title != "Usage of `"+c.Route.Name+"`" {
		t.Fatal("usage.Embed.Title")
	}
}

func TestFlagsUsageErr(t *testing.T) {
	_, _, _, _, _, _,
		c := ctxDummy("uwu")
	_, _,
		flags := flagsDummy(c)
	flags.err = fmt.Errorf("this is an error")

	usage := flags.Usage()

	if usage.Content != "**Error:** "+flags.err.Error() {
		t.Fatal("usage.Content")
	}
	if usage.Embed.Title != "Usage of `"+c.Route.Name+"`" {
		t.Fatal("usage.Embed.Title")
	}
}

func TestFlagsParse(t *testing.T) {
	_, _, _, _, _, _,
		c := ctxDummy("uwu")
	_, args, flags := flagsDummy(c)

	err := flags.Parse()
	if err != nil {
		t.Fatal(err)
	}

	fargs := flags.Args()

	if len(args) != len(fargs) {
		t.Fatalf(
			"len(args) == %d; len(fargs) == %d",
			len(args), len(fargs),
		)
	}

	for i := 0; i < len(args) && i < len(fargs); i++ {
		if args[i] != fargs[i] {
			t.Fatalf(
				"args[%d] == %q; fargs[%d] == %q",
				i, args[i],
				i, fargs[i],
			)
		}
	}
}

func TestFlagsParseErr(t *testing.T) {
	_, _, _, _, _, _,
		c := ctxDummy("uwu")
	_, args, flags := flagsDummy(c)
	args[0] = "-badflag"

	err := flags.Parse()
	if err == nil {
		t.Fatal("err is nil")
	}
}

func TestFlagsFlags(t *testing.T) {
	_, _, _, _, _, _,
		c := ctxDummy("uwu")
	_, args, flags := flagsDummy(c)
	args[0] = "-foo"
	args[1] = "bar"
	args[2] = "-baz"

	foo := flags.String("foo", "", "foo flag")
	baz := flags.Bool("baz", false, "baz flag")

	err := flags.Parse()
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
	_, _, _, _, _, _,
		c := ctxDummy("uwu")
	_, _, flags := flagsDummy(c)
	const oout = "hello world"

	_, err := flags.FlagSet.Output().Write([]byte(oout))
	if err != nil {
		t.Fatal(err)
	}

	out := flags.Output()
	if out != oout {
		t.Fatalf(
			"out == %q; oout == %q",
			out, oout,
		)
	}
}

type dummyWriter struct{}

func (d dummyWriter) Write(p []byte) (int, error) {
	return 0, nil
}

func TestFlagsOutputOther(t *testing.T) {
	_, _, _, _, _, _,
		c := ctxDummy("uwu")
	_, _, flags := flagsDummy(c)
	flags.SetOutput(dummyWriter{})
	const oout = "hello world"

	_, err := flags.FlagSet.Output().Write([]byte(oout))
	if err != nil {
		t.Fatal(err)
	}

	out := flags.Output()
	if out != "" {
		t.Fatalf("out != %q", "")
	}
}

type dummyValue struct{}

func (d dummyValue) String() string {
	return "dummy"
}
func (d dummyValue) Set(string) error {
	return fmt.Errorf("dummy")
}

func TestFlagsVisitor(t *testing.T) {
	_, _, _, _, _, _,
		c := ctxDummy("uwu")

	rep := c.Reply()
	rep.Embed = &dg.MessageEmbed{
		Title:       "Usage of `" + c.Route.Name + "`",
		Description: c.Route.Desc,
		Fields:      make([]*dg.MessageEmbedField, 0),
	}

	val := dummyValue{}
	visitor(rep)(&flag.Flag{
		Name:     "foo",
		Usage:    "do the foo thing idk",
		Value:    val,
		DefValue: val.String(),
	})
}
