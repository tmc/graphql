package parser

import (
	"errors"

	"github.com/tmc/graphql"
	"github.com/tmc/graphql/internal/parser"
)

var (
	// ErrMalformedQuery means there was a parsing error with the provided query
	ErrMalformedQuery = errors.New("parser: malformed graphql query")
)

// Parse attempts to parse a graphql.Call from a byte slice.
func Parse(query []byte) (graphql.Call, error) {
	result, err := parser.Parse("", query)
	if err != nil {
		return graphql.Call{}, ErrMalformedQuery
	}
	return result.(graphql.Call), nil

}
