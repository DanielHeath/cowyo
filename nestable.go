package main

import (
	"fmt"
	"strings"
	"unicode"
)

type NestableDoc struct {
	Children []*Nestable
}

func (d NestableDoc) String() string {
	strs := make([]string, len(d.Children))
	for i, c := range d.Children {
		strs[i] = c.String()
	}
	return strings.Join(strs, "\n")
}

type Nestable struct {
	Text     string
	Children []*Nestable
}

func ParseNestable(text string) (NestableDoc, error) {
	root := NewNestable()
	stack := []*Nestable{&root}

	prevLineLevel := -1 // impossible; only the root can have this value.
	for idx, line := range strings.Split(text, "\n") {
		depth := indentDepth(line)
		thisLine := NewNestable()
		thisLine.Text = line[depth:]
		if depth == prevLineLevel {
			// Close off the previous line
			fmt.Printf("shorten stack from: %d\n", len(stack))
			stack = stack[:len(stack)-1]
			fmt.Printf("shorten stack to: %d\n", len(stack))
		} else if depth == prevLineLevel+1 {
		} else if depth > prevLineLevel {
			return NestableDoc{}, fmt.Errorf("Line %d jumped from depth %d to depth %d", idx, prevLineLevel, depth)
		} else {
			stack = stack[:len(stack)-1+depth-prevLineLevel]
		}
		fmt.Printf("prev %d now %d for '%s'\n", prevLineLevel, depth, line)
		fmt.Printf("%#v\n", NestableDoc{root.Children}.String())
		fmt.Printf("stack depth: %d\n", len(stack))

		prevLineLevel = depth
		// add ourselves as a child
		stack[len(stack)-1].Children = append(stack[len(stack)-1].Children, &thisLine)
		// we're the deepest now
		stack = append(stack, &thisLine)
	}
	return NestableDoc{root.Children}, nil
}

func NewNestable() Nestable {
	return Nestable{Children: []*Nestable{}}
}

func (n Nestable) String() string {
	return n.IndentedString("")
}

func (n Nestable) IndentedString(indent string) string {
	result := []string{indent + n.Text}
	for _, c := range n.Children {
		result = append(result, c.IndentedString(indent+" "))
	}
	return strings.Join(result, "\n")
}

func indentDepth(s string) int {
	var idx int
	var r rune
	for idx, r = range s {
		if !unicode.IsSpace(r) {
			break
		}
	}
	return idx
}
