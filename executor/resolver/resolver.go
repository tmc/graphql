package resolver

import "github.com/tmc/graphql"

type Resolver interface {
	Resolve(interface{}, *graphql.Field) (interface{}, error)
}
