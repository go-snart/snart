package route_test

import (
	"testing"

	"github.com/go-snart/snart/route"
)

func TestSplitNormal(t *testing.T) {
	args := route.Split("hello world")

	argsTest := []string{"hello", "world"}

	if len(args) != len(argsTest) {
		t.Fatalf(
			"len(args) == %d != len(argsTest) == %d",
			len(args), len(argsTest),
		)
	}

	for i, argTest := range argsTest {
		if args[i] != argTest {
			t.Fatalf(
				"(at %d) expected %q, found %q",
				i, argTest, args[i],
			)
		}
	}
}

func TestSplitSingleQuote(t *testing.T) {
	args := route.Split("hello `foo bar` world")

	argsTest := []string{"hello", "foo bar", "world"}

	if len(args) != len(argsTest) {
		t.Fatalf(
			"len(args) == %d != len(argsTest) == %d",
			len(args), len(argsTest),
		)
	}

	for i, argTest := range argsTest {
		if args[i] != argTest {
			t.Fatalf(
				"(at %d) expected %q, found %q",
				i, argTest, args[i],
			)
		}
	}
}

func TestSplitDoubleQuote(t *testing.T) {
	args := route.Split("hello ``foo ` bar`` world")

	argsTest := []string{"hello", "foo ` bar", "world"}

	if len(args) != len(argsTest) {
		t.Fatalf(
			"len(args) == %d != len(argsTest) == %d",
			len(args), len(argsTest),
		)
	}

	for i, argTest := range argsTest {
		if args[i] != argTest {
			t.Fatalf(
				"(at %d) expected %q, found %q",
				i, argTest, args[i],
			)
		}
	}
}
