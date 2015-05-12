package parser_test

import (
	"encoding/json"
	"fmt"

	"github.com/tmc/graphql/parser"
)

const exampleQuery = `{
  user(id: 42) {
	friends(isViewerFriend: true, first: 10) {
	  nodes {
		name
	  }
	}
  }
}`

func ExampleParse() {
	result, err := parser.ParseOperation([]byte(exampleQuery))
	if err != nil {
		fmt.Println("err:", err)
	}
	asjson, _ := json.MarshalIndent(result, "", " ")
	fmt.Println(string(asjson))
	// output:
	// {
	//  "Name": "node",
	//  "Arguments": [
	//   {
	//    "Name": "id",
	//    "Value": "42"
	//   }
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
