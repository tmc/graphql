package schema_test

import (
	"encoding/json"
	"fmt"

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
