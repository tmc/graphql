// Program basic_graphql_server shows a simple HTTP server that exposes a bare schema.
//
// Example:
//  $ go get github.com/tmc/graphql/example/basic_graphql_server
//  $ basic_graphql_server &
//  $ curl -g 'http://localhost:8080/?q={__schema__{root_fields{name,description}}}'
//  {"data":[{"root_fields":[{"description": "Schema entry root field","name":"__schema__"}]}}]
//
// Here we see the server showing the available root fields ("schema").
package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/tmc/graphql"
	"github.com/tmc/graphql/executor"
	"github.com/tmc/graphql/executor/resolver"
	"github.com/tmc/graphql/handler"
	"github.com/tmc/graphql/schema"
	"golang.org/x/net/context"
)

var listenAddr = flag.String("l", ":8080", "listen addr")

type nowProvider struct {
	start time.Time
}

func (n *nowProvider) now(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
	return time.Now(), nil
}

func (n *nowProvider) uptime(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
	return time.Now().Sub(n.start).Seconds(), nil
}

func (n *nowProvider) GraphQLTypeInfo() schema.GraphQLTypeInfo {
	return schema.GraphQLTypeInfo{
		Name:        "nowProvider",
		Description: "example root field provider",
		Fields: map[string]*schema.GraphQLFieldSpec{
			"now":    {"now", "Provides the current server time", n.now, []graphql.Argument{}, true},
			"uptime": {"uptime", "Provides the current server uptime", n.uptime, []graphql.Argument{}, true},
		},
	}
}

func main() {
	flag.Parse()
	// create a new schema (which self-registers)
	now := &nowProvider{time.Now()}

	schema := schema.New()
	schema.Register(now)

	executor := executor.New(schema)
	handler := handler.New(executor)
	mux := http.NewServeMux()
	mux.Handle("/", handler)
	log.Fatalln(http.ListenAndServe(*listenAddr, mux))
}
