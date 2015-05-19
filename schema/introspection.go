package schema

import (
	"log"

	"github.com/tmc/graphql"
	"github.com/tmc/graphql/executor/resolver"
)

type GraphQLTypeIntrospector struct {
	typeInfo GraphQLTypeInfo
	//schema   *Schema
}

func newIntrospectionField(typeInfo GraphQLTypeInfo) GraphQLFieldFunc {
	return func(r resolver.Resolver, f *graphql.Field) (interface{}, error) {
		log.Println("creating introspector for", typeInfo)
		return &GraphQLTypeIntrospector{
			typeInfo: typeInfo,
			//schema:   schema,
		}, nil
	}
}

func (i *GraphQLTypeIntrospector) GraphQLTypeInfo() GraphQLTypeInfo {
	return WithIntrospectionField(GraphQLTypeInfo{
		Name:        "GraphQLTypeIntrospector",
		Description: "Provides type introspection capabilities",
		Fields: map[string]*GraphQLFieldSpec{
			"name": &GraphQLFieldSpec{
				Name:        "name",
				Description: "Returns the name of the GraphQL type.",
				Func:        i.name,
			},
			"fields": &GraphQLFieldSpec{
				Name:        "fields",
				Description: "Returns the fields present on a GraphQL type.",
				Func:        i.fields,
			},
		},
	})
}

func (i *GraphQLTypeIntrospector) name(r resolver.Resolver, f *graphql.Field) (interface{}, error) {
	return i.typeInfo.Name, nil
}

func (i *GraphQLTypeIntrospector) fields(r resolver.Resolver, f *graphql.Field) (interface{}, error) {
	result := []interface{}{}
	for key := range i.typeInfo.Fields {
		res, err := r.Resolve(i.typeInfo.Fields[key], f)
		if err != nil {
			return nil, err
		}
		result = append(result, res)
	}
	return result, nil
}
