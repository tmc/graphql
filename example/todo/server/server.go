package server

import (
	"log"
	"net/http"

	"github.com/tmc/graphql/example/todo/graph"
	"github.com/tmc/graphql/executor"
	"github.com/tmc/graphql/handler"
	"github.com/tmc/graphql/schema"
)

type App struct {
	address string
}

var Application *App

func NewApp(address string) *App {
	app := new(App)
	app.address = address

	return app
}

func (app *App) RunServer() {
	g := graph.NewGraph()

	schema := schema.New()
	schema.Register(g)

	executor := executor.New(schema)
	handler := handler.New(executor)
	mux := http.NewServeMux()
	mux.Handle("/", handler)
	log.Fatalln(http.ListenAndServe(app.address, mux))
}
