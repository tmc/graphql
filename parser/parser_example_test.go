package parser_test

import (
	"fmt"

	"github.com/tmc/graphql/parser"
)

const exampleQuery = `node(42){answer}`

func ExampleParse() {
	result, err := parser.Parse("example.graphql", []byte("foobar"))
	fmt.Println(result, err)
	// output:
	// foobar <nil>
}
