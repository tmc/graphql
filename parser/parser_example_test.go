package parser_test

import (
	"encoding/json"
	"fmt"

	"github.com/tmc/graphql/parser"
)

const exampleQuery = `node(42){id,answer,{towel,planet}}`

func ExampleParse() {
	result, err := parser.Parse([]byte(exampleQuery))
	if err != nil {
		fmt.Println("err:", err)
	}
	asjson, _ := json.MarshalIndent(result, "", " ")
	fmt.Println(string(asjson))
	// output:
	//  {
	//  "Name": "node",
	//  "Arguments": [
	//   "42"
	//  ],
	//  "Fields": [
	//   {
	//    "Name": "id"
	//   },
	//   {
	//    "Name": "answer"
	//   },
	//   {
	//    "Fields": [
	//     {
	//      "Name": "towel"
	//     },
	//     {
	//      "Name": "planet"
	//     }
	//    ]
	//   }
	//  ]
	// }
}
