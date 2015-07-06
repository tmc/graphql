package parser_test

import (
	"testing"

	"github.com/tmc/graphql/parser"
)

func TestMalformedQuery(t *testing.T) {
	op, err := parser.ParseOperation(nil)
	if !parser.IsMalformedOperation(err) {
		t.Error("Expected malformed operation")
	}
	if op != nil {
		t.Error("Expected nil result")
	}
	if es := err.Error(); es != "parser: malformed graphql operation: 1:1 (0): no match found" {
		t.Errorf("got unexpected error: '%v'", es)
	}
}

func TestMultipleOperations(t *testing.T) {
	multi := `
	{foo}
	{bar}
	`
	op, err := parser.ParseOperation([]byte(multi))
	if err != parser.ErrMultipleOperations {
		t.Error("Expected multiple operations error")
	}
	if op != nil {
		t.Error("Expected nil result")
	}
}
