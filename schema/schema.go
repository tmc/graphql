package schema

import (
	"sort"

	"github.com/tmc/graphql"
	"github.com/tmc/graphql/executor/resolver"
	"golang.org/x/net/context"
)

// Schema represents the registered types that know how to respond to root fields.
type Schema struct {
	registeredTypes map[string]GraphQLTypeInfo
	rootFields      map[string]*GraphQLFieldSpec
}

// New prepares a new Schema.
func New() *Schema {
	s := &Schema{
		registeredTypes: map[string]GraphQLTypeInfo{},
		rootFields:      map[string]*GraphQLFieldSpec{},
	}
	// self-register
	s.Register(s)
	return s
}

// Register registers a new type
func (s *Schema) Register(t GraphQLType) {
	typeInfo := t.GraphQLTypeInfo()
	s.registeredTypes[t.GraphQLTypeInfo().Name] = typeInfo
	// TODO(tmc): collision handling
	for name, fieldSpec := range typeInfo.Fields {
		if fieldSpec.IsRoot {
			s.rootFields[name] = fieldSpec
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

func (s *Schema) RootFields() map[string]*GraphQLFieldSpec {
	return s.rootFields
}

func (s *Schema) GetTypeInfo(o GraphQLType) GraphQLTypeInfo {
	return s.registeredTypes[o.GraphQLTypeInfo().Name]
}

func (s *Schema) RegisteredTypes() map[string]GraphQLTypeInfo {
	return s.registeredTypes
}

// The below makes Schema itsself a GraphQLType and provides the root field of 'schema'

func (s *Schema) GraphQLTypeInfo() GraphQLTypeInfo {
	return GraphQLTypeInfo{
		Name:        "Schema",
		Description: "Root schema object",
		Fields: map[string]*GraphQLFieldSpec{
			"__schema":    {"__schema", "Schema entry root field", s.handleSchemaCall, nil, true},
			"__types":     {"__types", "Introspection of registered types", s.handleTypesCall, nil, true},
			"root_fields": {"root_fields", "List fields that are exposed at the root of the GraphQL schema.", s.handleRootFields, nil, false},
		},
	}
}

func (s *Schema) handleSchemaCall(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
	return s, nil
}

func (s *Schema) handleTypesCall(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
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

func (s *Schema) handleRootFields(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
	rootFields := []string{}
	for rootField := range s.rootFields {
		rootFields = append(rootFields, rootField)
	}
	sort.Strings(rootFields)
	result := make([]*GraphQLFieldSpec, 0, len(rootFields))
	for _, field := range rootFields {
		result = append(result, s.rootFields[field])
	}
	return result, nil
}
