package parser

import (
	"errors"
	"log"

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
func ParseOperation(query []byte) (*graphql.Operation, error) {
	result, err := parser.Parse("", query)
	if err != nil {
		log.Println("parse error:", err)
		return nil, ErrMalformedOperation
	}
	doc, ok := result.(graphql.Document)
	if !ok {
		return nil, ErrMalformedOperation
	}
	switch len(doc.Operations) {
	case 1:
		return &doc.Operations[0], nil
	case 0:
		return nil, ErrMalformedOperation
	default:
		return nil, ErrMultipleOperations
	}
}
