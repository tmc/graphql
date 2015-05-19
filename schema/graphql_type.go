package schema

import (
	"github.com/tmc/graphql"
	"github.com/tmc/graphql/executor/resolver"
)

// GraphQLFieldFunc is the type that can generate a response from a graphql.Field.
type GraphQLFieldFunc func(resolver.Resolver, *graphql.Field) (interface{}, error)

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

type Scalar struct {
	Value interface{}
}

func (s Scalar) GraphQLTypeInfo() GraphQLTypeInfo {
	return GraphQLTypeInfo{
		Name:        "Scalar",
		Description: "A scalar value",
	}
}

type GraphQLFieldSpec struct {
	Name        string
	Description string
	Func        GraphQLFieldFunc
	Arguments   []graphql.Argument
	IsRootCall  bool
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

func (g *GraphQLFieldSpec) name(r resolver.Resolver, f *graphql.Field) (interface{}, error) {
	return g.Name, nil
}

func (g *GraphQLFieldSpec) description(r resolver.Resolver, f *graphql.Field) (interface{}, error) {
	return g.Description, nil
}

type GraphQLFieldSpecMap map[string]*GraphQLFieldSpec

func (g GraphQLFieldSpecMap) GraphQLTypeInfo() GraphQLTypeInfo {
	return GraphQLTypeInfo{
		Name:        "GraphQLFieldSpecMap",
		Description: "A collection of GraphQLFieldSpec objects",
	}
}
