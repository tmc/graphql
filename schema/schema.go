package schema

import (
	"fmt"
	"sort"

	"github.com/tmc/graphql"
)

// Schema represents the registered types that know how to respond to root calls.
type Schema struct {
	registeredTypes []GraphQLType
	rootCalls       map[string]FieldHandler
}

// FieldHandler is the type that can generate a response from a graphql.Field.
type FieldHandler func(*graphql.Field) (interface{}, error)

// GraphQLType is the interface that root call providers conform to.
type GraphQLType interface {
	RootCalls() map[string]FieldHandler
}

// New prepares a new Schema.
func New() *Schema {
	s := &Schema{
		registeredTypes: []GraphQLType{},
		rootCalls:       map[string]FieldHandler{},
	}
	// self-register
	s.Register(s)
	return s
}

// HandleField dispatches a graphql.Field to the appropriate registered type.
func (s *Schema) HandleField(f *graphql.Field) (interface{}, error) {
	handler, ok := s.rootCalls[f.Name]
	if !ok {
		return nil, fmt.Errorf("schema: no registered types handle the root call '%s'", f.Name)
	}
	return handler(f)
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
func (s *Schema) RootCalls() map[string]FieldHandler {
	return map[string]FieldHandler{
		"schema": s.handleSchemaCall,
	}
}

func (s *Schema) handleSchemaCall(f *graphql.Field) (interface{}, error) {
	result := map[string]interface{}{}
	rootCalls := []string{}
	for _, selection := range f.Selections {
		if selection.Field.Name == "root_calls" {
			for rootCall := range s.rootCalls {
				rootCalls = append(rootCalls, rootCall)
			}
			sort.Strings(rootCalls)
			result["root_calls"] = rootCalls
		}
	}
	return result, nil
}
