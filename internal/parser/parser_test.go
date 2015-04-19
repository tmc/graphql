package parser_test

import (
	"testing"

	"github.com/tmc/graphql/internal/parser"
)

var shouldParse = []string{
	`foo()`,
	`node(42)`,
	`foo(){id}`,
	`foo(1,2){id}`,
	`foo(1,2){id,{nest,some,{fields}}}`,
	`node(42){id,friends.top(10)}`,
}

func TestSuccessfulParses(t *testing.T) {
	for i, in := range shouldParse {
		_, err := parser.Parse("parser_test.go", []byte(in))
		if err != nil {
			t.Errorf("case %d: %v", i+1, err)
		}
	}
}
