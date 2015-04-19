package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/tmc/graphql/parser"
	"github.com/tmc/graphql/schema"
)

// Error represents an error the occured while parsing a graphql query or while generating a response.
type Error struct {
	Message string `json:"message"`
}

// Result represents a graphql query result.
type Result struct {
	Data  interface{} `json:"data,omitempty"`
	Error *Error      `json:"error,omitempty"`
}

// SchemaHandler makes a schema.Schema querable via HTTP
type SchemaHandler struct {
	schema *schema.Schema
}

// New constructs a SchemaHandler from a schema.
func New(schema *schema.Schema) *SchemaHandler {
	return &SchemaHandler{schema: schema}
}

func writeErr(w io.Writer, err error) {
	writeJSON(w, Result{Error: &Error{Message: err.Error()}})
}
func writeJSON(w io.Writer, data interface{}) {
	json.NewEncoder(w).Encode(data)
}

// ServeHTTP provides an entrypoint into a graphql schema. It pulls the query from
// the 'q' GET parameter.
func (h *SchemaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//TODO(tmc): reject non-GET requests
	q := r.URL.Query().Get("q")
	call, err := parser.Parse([]byte(q))
	if err != nil {
		writeErr(w, err)
		return
	}
	result, err := h.schema.HandleCall(call)
	if err != nil {
		writeErr(w, err)
		return
	}
	writeJSON(w, Result{Data: result})
}
