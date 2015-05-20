package schema

import (
	"sort"

	"github.com/tmc/graphql"
	"github.com/tmc/graphql/executor/resolver"
)

// Schema represents the registered types that know how to respond to root calls.
type Schema struct {
	registeredTypes map[string]GraphQLTypeInfo
	rootCalls       map[string]*GraphQLFieldSpec
}

// New prepares a new Schema.
func New() *Schema {
	s := &Schema{
		registeredTypes: map[string]GraphQLTypeInfo{},
		rootCalls:       map[string]*GraphQLFieldSpec{},
	}
	// self-register
	s.Register(s)
	// register special introspection type
	//i := &GraphQLTypeIntrospector{schema: s}
	i := &GraphQLTypeIntrospector{}
	s.Register(i)
	//s.Register(&GraphQLFieldSpec{})
	return s
}

// Register registers a new type
func (s *Schema) Register(t GraphQLType) {
	typeInfo := t.GraphQLTypeInfo()
	s.registeredTypes[t.GraphQLTypeInfo().Name] = typeInfo
	// TODO(tmc): collision handling
	for name, fieldSpec := range typeInfo.Fields {
		if fieldSpec.IsRootCall {
			s.rootCalls[name] = fieldSpec
		}
	}
}

func WithIntrospectionField(typeInfo GraphQLTypeInfo) GraphQLTypeInfo {
	introSpectionFunc := newIntrospectionField(typeInfo)
	typeInfo.Fields["__type__"] = &GraphQLFieldSpec{
		Name:        "__type__",
		Description: "Introspection field that exposes field and type information",
		Func:        introSpectionFunc,
	}
	return typeInfo
}

// External entrypoint

/*
// HandleField dispatches a graphql.Field to the appropriate registered type.
func (s *Schema) HandleField(f *graphql.Field) (interface{}, error) {
	handler, ok := s.rootCalls[f.Name]
	if !ok {
		return nil, fmt.Errorf("schema: no registered types handle the root call '%s'", f.Name)
	}
	return handler.Func(f)
}
*/

func (s *Schema) RootCalls() map[string]*GraphQLFieldSpec {
	return s.rootCalls
}

func (s *Schema) GetTypeInfo(o GraphQLType) GraphQLTypeInfo {
	panic(s)
	return s.registeredTypes[o.GraphQLTypeInfo().Name]
}

func (s *Schema) RegisteredTypes() map[string]GraphQLTypeInfo {
	return s.registeredTypes
}

// The below makes Schema itsself a GraphQLType and provides the root entry call of 'schema'

func (s *Schema) GraphQLTypeInfo() GraphQLTypeInfo {
	return GraphQLTypeInfo{
		Name:        "Schema",
		Description: "Root schema object",
		Fields: map[string]*GraphQLFieldSpec{
			"schema":     {"schema", "Schema entry root call", s.handleSchemaCall, nil, true},
			"types":      {"types", "Introspection of registered types", s.handleTypesCall, nil, true},
			"root_calls": {"root_calls", "List root_calls of registered types", s.handleRootCalls, nil, false},
		},
	}
}

func (s *Schema) handleSchemaCall(r resolver.Resolver, f *graphql.Field) (interface{}, error) {
	return s, nil
}

func (s *Schema) handleTypesCall(r resolver.Resolver, f *graphql.Field) (interface{}, error) {
	typeNames := make([]string, 0, len(s.registeredTypes))
	for typeName := range s.registeredTypes {
		typeNames = append(typeNames, typeName)
	}
	sort.Strings(typeNames)
	result := make([]GraphQLTypeInfo, 0, len(typeNames))
	for _, typeName := range typeNames {
		result = append(result, s.registeredTypes[typeName])
	}
	return result, nil
}

func (s *Schema) handleRootCalls(r resolver.Resolver, f *graphql.Field) (interface{}, error) {
	rootCalls := []string{}
	for rootCall := range s.rootCalls {
		rootCalls = append(rootCalls, rootCall)
	}
	sort.Strings(rootCalls)
	return rootCalls, nil
}
