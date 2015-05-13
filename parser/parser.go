package parser

import (
	"errors"

	"github.com/tmc/graphql"
	"github.com/tmc/graphql/internal/parser"
)

var (
	// ErrMalformedOperation means there was a parsing error with the provided query
	ErrMalformedOperation = errors.New("parser: malformed graphql operation")
	// ErrMultipleOperations
	ErrMultipleOperations = errors.New("parser: multiple graphql operations")
)

// ParseOperation attempts to parse a graphql.Operation from a byte slice.
func ParseOperation(query []byte) (graphql.Operation, error) {
	result, err := parser.Parse("", query)
	if err != nil {
		panic(err)
		// TODO(tmc))
		return graphql.Operation{}, ErrMalformedOperation
	}
	doc, ok := result.(graphql.Document)
	if !ok {
		return graphql.Operation{}, ErrMalformedOperation
	}
	switch len(doc.Operations) {
	case 1:
		return doc.Operations[0], nil
	case 0:
		return graphql.Operation{}, ErrMalformedOperation
	default:
		return graphql.Operation{}, ErrMultipleOperations
	}
}
