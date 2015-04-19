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
