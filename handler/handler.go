package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/tmc/graphql/executor"
	"github.com/tmc/graphql/parser"
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

// ExecutorHandler makes a executor.Executor querable via HTTP
type ExecutorHandler struct {
	executor *executor.Executor
}

// New constructs a ExecutorHandler from a executor.
func New(executor *executor.Executor) *ExecutorHandler {
	return &ExecutorHandler{executor: executor}
}

func writeErr(w io.Writer, err error) {
	writeJSON(w, Result{Error: &Error{Message: err.Error()}})
}
func writeJSON(w io.Writer, data interface{}) {
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Println("error writing json response:", err)
		// attempt to write error
		writeErr(w, err)
	}
}

// ServeHTTP provides an entrypoint into a graphql executor. It pulls the query from
// the 'q' GET parameter.
func (h *ExecutorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//TODO(tmc): reject non-GET requests
	q := r.URL.Query().Get("q")
	log.Println("query:", q)
	operation, err := parser.ParseOperation([]byte(q))
	if err != nil {
		log.Println("error parsing:", err)
		writeErr(w, err)
		return
	}
	// if err := h.validator.Validate(operation); err != nil { writeErr(w, err); return }
	result, err := h.executor.HandleOperation(operation)
	if err != nil {
		writeErr(w, err)
	} else {
		writeJSON(w, Result{Data: result})
	}
}