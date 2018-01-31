package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestRoundTripNestable(t *testing.T) {
	cases := []string{
		"\n",
		"",
		"Foo",
		strings.TrimSpace(`
Foo
Bar
    `),
		strings.TrimSpace(`
Foo
 Bar
  Baz
    `),
		strings.TrimSpace(`
Foo
 Bar
Baz
 Bim

 Foo (child of an empty parent)
Bens
 Bent
  Broom
   Breaks
    Badly
  Bims
   Bent
Broom
Breaks
    `),
	}
	for _, v := range cases {
		parsed, err := ParseNestable(v)
		if err != nil {
			panic(err)
		}
		if v != parsed.String() {
			panic(fmt.Sprintf("Expected \n%#v$\nto round-trip to itself; got\n%#v$\nfrom '%#v'", v, parsed.String(), parsed))
		}
	}
}
