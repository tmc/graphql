package parser_test

import (
	"encoding/json"
	"fmt"

	"github.com/tmc/graphql/parser"
)

const exampleQuery = `{ foo { field } }`

func ExampleParse() {
	result, err := parser.ParseOperation([]byte(exampleQuery))
	if err != nil {
		fmt.Println("err:", err)
	} else {
		asjson, _ := json.MarshalIndent(result, "", " ")
		fmt.Println(string(asjson))
	}
	// output:
	// {
	//  "Type": "query",
	//  "Selections": [
	//   {
	//    "Field": {
	//     "Name": "foo",
	//     "Selections": [
	//      {
	//       "Field": {
	//        "Name": "field"
	//       }
	//      }
	//     ]
	//    }
	//   }
	//  ]
	// }
}
