package schema

import (
	"fmt"
	"sort"

	"github.com/tmc/graphql"
)

// Schema represents the registered types that know how to respond to root calls.
type Schema struct {
	registeredTypes []GraphQLType
	rootCalls       map[string]CallHandler
}

// CallHandler is the type that can generate a response from a graphql Operation.
type CallHandler func(graphql.Operation) (interface{}, error)

// GraphQLType is the interface that root call providers conform to.
type GraphQLType interface {
	RootCalls() map[string]CallHandler
}

// New prepares a new Schema.
func New() *Schema {
	s := &Schema{
		registeredTypes: []GraphQLType{},
		rootCalls:       map[string]CallHandler{},
	}
	// self-register
	s.Register(s)
	return s
}

// HandleCall dispatches a graphql.Operation to the appropriate registered type.
func (s *Schema) HandleCall(c graphql.Operation) (interface{}, error) {
	handler, ok := s.rootCalls[c.Name]
	if !ok {
		return nil, fmt.Errorf("schema: no registered types handle the root call '%s'", c.Name)
	}
	return handler(c)
}

// Register registers a new type that provides root calls.
func (s *Schema) Register(t GraphQLType) {
	s.registeredTypes = append(s.registeredTypes, t)
	for call, handler := range t.RootCalls() {
		// TODO(tmc): collision handling
		s.rootCalls[call] = handler
	}
}

// RootCalls returns the root call that Schema itself provides.
func (s *Schema) RootCalls() map[string]CallHandler {
	return map[string]CallHandler{
		"schema": s.handleSchemaCall,
	}
}

func (s *Schema) handleSchemaCall(c graphql.Operation) (interface{}, error) {
	result := map[string]interface{}{}
	rootCalls := []string{}
	for _, selection := range c.Selections {
		if selection.FieldName == "root_calls" {
			for rootCall := range s.rootCalls {
				rootCalls = append(rootCalls, rootCall)
			}
			sort.Strings(rootCalls)
			result["root_calls"] = rootCalls
		}
	}
	return result, nil
}
