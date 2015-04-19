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
	schema := schema.New()

	mux := http.NewServeMux()
	mux.Handle("/", handler.New(schema))
	log.Fatalln(http.ListenAndServe(*listenAddr, mux))
}
