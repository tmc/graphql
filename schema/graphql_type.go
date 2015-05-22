package schema

import (
	"sort"

	"github.com/tmc/graphql"
	"github.com/tmc/graphql/executor/resolver"
	"golang.org/x/net/context"
)

// GraphQLFieldFunc is the type that can generate a response from a graphql.Field.
type GraphQLFieldFunc func(context.Context, resolver.Resolver, *graphql.Field) (interface{}, error)

type GraphQLTypeInfo struct {
	Name        string
	Description string
	Fields      GraphQLFieldSpecMap
	//Init        func(map[string]interface{}) (GraphQLType, error)
}

// GraphQLType is the interface that all GraphQL types satisfy
type GraphQLType interface {
	GraphQLTypeInfo() GraphQLTypeInfo
}

func (g GraphQLTypeInfo) GraphQLTypeInfo() GraphQLTypeInfo {
	return GraphQLTypeInfo{
		Name:        "GraphQLTypeInfo",
		Description: "Holds information about GraphQLTypeInfo",
		Fields: GraphQLFieldSpecMap{
			"name": {
				Name:        "name",
				Description: "The name of the type.",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					return r.Resolve(ctx, g.Name, f)
				},
			},
			"description": {
				Name:        "description",
				Description: "The description of the type.",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					return r.Resolve(ctx, g.Description, f)
				},
			},
			"fields": {
				Name:        "fields",
				Description: "The fields associated with the type.",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					fields := make([]string, 0, len(g.Fields))
					for fieldName := range g.Fields {
						fields = append(fields, fieldName)
					}
					sort.Strings(fields)
					result := make([]*GraphQLFieldSpec, 0, len(g.Fields))
					for _, fieldName := range fields {
						result = append(result, g.Fields[fieldName])
					}
					return r.Resolve(ctx, result, f)
				},
			},
		},
	}
}

type Scalar struct {
	Value interface{}
}

func (s Scalar) GraphQLTypeInfo() GraphQLTypeInfo {
	return GraphQLTypeInfo{
		Name:        "Scalar",
		Description: "A scalar value",
	}
}

// GraphQLFieldSpec describes a field associated with a type in a GraphQL schema.
type GraphQLFieldSpec struct {
	Name        string
	Description string
	Func        GraphQLFieldFunc
	Arguments   []graphql.Argument // Describes any arguments the field accepts
	IsRoot      bool               // If true, this field should be exposed at the root of the GraphQL schema
}

func (g *GraphQLFieldSpec) GraphQLTypeInfo() GraphQLTypeInfo {
	return GraphQLTypeInfo{
		Name:        "GraphQLFieldSpec",
		Description: "A GraphQL field specification",
		Fields: map[string]*GraphQLFieldSpec{
			"name":        {"name", "Field name", g.name, nil, false},
			"description": {"description", "Field description", g.description, nil, false},
		},
	}
}

func (g *GraphQLFieldSpec) name(_ context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
	return g.Name, nil
}

func (g *GraphQLFieldSpec) description(_ context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
	return g.Description, nil
}

type GraphQLFieldSpecMap map[string]*GraphQLFieldSpec

func (g GraphQLFieldSpecMap) GraphQLTypeInfo() GraphQLTypeInfo {
	return GraphQLTypeInfo{
		Name:        "GraphQLFieldSpecMap",
		Description: "A collection of GraphQLFieldSpec objects",
	}
}
