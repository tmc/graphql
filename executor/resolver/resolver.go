package resolver

import (
	"github.com/tmc/graphql"
	"golang.org/x/net/context"
)

type Resolver interface {
	Resolve(context.Context, interface{}, *graphql.Field) (interface{}, error)
}
