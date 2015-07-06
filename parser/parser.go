package parser

import (
	"errors"
	"fmt"

	"github.com/tmc/graphql"
	"github.com/tmc/graphql/internal/parser"
)

var (
	// ErrMultipleOperations
	ErrMultipleOperations = errors.New("parser: multiple graphql operations")
)

// ErrMalformedOperation means there was a parsing error with the provided query
type ErrMalformedOperation struct {
	underlying error
}

func IsMalformedOperation(err error) bool {
	_, ok := err.(ErrMalformedOperation)
	return ok
}

func (e ErrMalformedOperation) Error() string {
	return fmt.Sprintf("parser: malformed graphql operation: %v", e.underlying)
}

// ParseOperation attempts to parse a graphql.Operation from a byte slice.
func ParseOperation(query []byte) (*graphql.Operation, error) {
	result, err := parser.Parse("", query)
	if err != nil {
		return nil, ErrMalformedOperation{err}
	}
	doc, ok := result.(graphql.Document)
	if !ok {
		return nil, ErrMalformedOperation{err}
	}
	switch len(doc.Operations) {
	case 1:
		return &doc.Operations[0], nil
	case 0:
		return nil, ErrMalformedOperation{fmt.Errorf("no operations")}
	default:
		return nil, ErrMultipleOperations
	}
}
