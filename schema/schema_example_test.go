package schema_test

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/tmc/graphql"
	"github.com/tmc/graphql/parser"
	"github.com/tmc/graphql/schema"
)

func ExampleSchema() {
	s := schema.New()
	call, err := parser.Parse("schema_test.go", []byte(`schema(){root_calls}`))
	if err != nil {
		fmt.Println(err)
	}
	result, err := s.HandleCall(call.(graphql.Call))
	if err != nil {
		fmt.Println(err)
	}
	asjson, _ := json.MarshalIndent(result, "", " ")
	fmt.Println(string(asjson))
	// output:
	// {
	//  "root_calls": [
	//   "schema"
	//  ]
	// }
}

type nowProvider struct{}

func (n *nowProvider) now(c graphql.Call) (interface{}, error) {
	return time.Now(), nil
}

func (n *nowProvider) RootCalls() map[string]schema.CallHandler {
	return map[string]schema.CallHandler{
		"now": n.now,
	}
}

func ExampleSchemaCustomType() {
	s := schema.New()
	s.Register(new(nowProvider))
	call, err := parser.Parse("schema_test.go", []byte(`schema(){root_calls}`))
	if err != nil {
		fmt.Println(err)
	}
	result, err := s.HandleCall(call.(graphql.Call))
	if err != nil {
		fmt.Println(err)
	}
	asjson, _ := json.MarshalIndent(result, "", " ")
	fmt.Println(string(asjson))
	// output:
	// {
	//  "root_calls": [
	//   "now",
	//   "schema"
	//  ]
	// }
}
