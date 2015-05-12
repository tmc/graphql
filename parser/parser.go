package parser

import (
	"errors"

	"github.com/tmc/graphql"
	"github.com/tmc/graphql/internal/parser"
)

var (
	// ErrMalformedOperation means there was a parsing error with the provided query
	ErrMalformedOperation = errors.New("parser: malformed graphql operation")
)

// ParseOperation attempts to parse a graphql.Operation from a byte slice.
func ParseOperation(query []byte) (graphql.Operation, error) {
	result, err := parser.Parse("", query)
	if err != nil {
		return graphql.Operation{}, ErrMalformedOperation
	}
	return result.(graphql.Operation), nil

}
