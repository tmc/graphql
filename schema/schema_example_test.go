package schema_test

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/tmc/graphql"
	"github.com/tmc/graphql/executor"
	"github.com/tmc/graphql/executor/resolver"
	"github.com/tmc/graphql/parser"
	"github.com/tmc/graphql/schema"
)

func ExampleSchema() {
	s := schema.New()
	call, err := parser.ParseOperation([]byte(`{__schema{root_fields{name}}}`))
	if err != nil {
		fmt.Println(err)
	}
	executor := executor.New(s)
	result, err := executor.HandleOperation(call)
	if err != nil {
		fmt.Println(err)
	}
	asjson, _ := json.MarshalIndent(result, "", " ")
	fmt.Println(string(asjson))
	// output:
	// [
	//  {
	//   "root_fields": [
	//    {
	//     "name": "__schema"
	//    },
	//    {
	//     "name": "__types"
	//    }
	//   ]
	//  }
	// ]
}

type nowProvider struct{}

func (n *nowProvider) now(r resolver.Resolver, f *graphql.Field) (interface{}, error) {
	return time.Now(), nil
}

func (n *nowProvider) GraphQLTypeInfo() schema.GraphQLTypeInfo {
	return schema.GraphQLTypeInfo{
		Name:        "now Provider",
		Description: "example root field provider",
		Fields: map[string]*schema.GraphQLFieldSpec{
			"now": {"now", "Provides the current server time", n.now, []graphql.Argument{}, true},
		},
	}
}
func ExampleSchemaCustomType() {
	s := schema.New()
	s.Register(new(nowProvider))
	call, err := parser.ParseOperation([]byte(`{__schema{root_fields{name}}}`))
	if err != nil {
		fmt.Println(err)
	}
	executor := executor.New(s)
	result, err := executor.HandleOperation(call)
	if err != nil {
		fmt.Println(err)
	}
	asjson, _ := json.MarshalIndent(result, "", " ")
	fmt.Println(string(asjson))
	// output:
	// [
	//  {
	//   "root_fields": [
	//    {
	//     "name": "__schema"
	//    },
	//    {
	//     "name": "__types"
	//    },
	//    {
	//     "name": "now"
	//    }
	//   ]
	//  }
	// ]
}
