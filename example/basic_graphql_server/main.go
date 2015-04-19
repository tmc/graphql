// Program basic_graphql_server shows a simple HTTP server that exposes a bare schema.
//
// Example:
//  $ go get github.com/tmc/graphql/example/basic_graphql_server
//  $ basic_graphql_server &
//  $ curl 'http://localhost:8080/?q=schema()\{root_calls\}'
//  {"data":{"root_calls":["schema"]}}
//
// Here we see the server showing the available root calls ("schema").
package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/tmc/graphql/schema"
	"github.com/tmc/graphql/schema/handler"
)

var listenAddr = flag.String("l", ":8080", "listen addr")

func main() {
	// create a new schema (which self-registers)
	schema := schema.New()
	mux := http.NewServeMux()
	mux.Handle("/", handler.New(schema))
	log.Fatalln(http.ListenAndServe(*listenAddr, mux))
}
