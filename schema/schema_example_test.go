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
	"golang.org/x/net/context"
)

func ExampleSchema() {
	s := schema.New()
	call, err := parser.ParseOperation([]byte(`{__schema__{root_fields{name}}}`))
	if err != nil {
		fmt.Println(err)
	}
	executor := executor.New(s)
	result, err := executor.HandleOperation(context.Background(), call)
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
	//     "name": "__schema__"
	//    },
	//    {
	//     "name": "__type"
	//    }
	//   ]
	//  }
	// ]
}

type nowProvider struct{}

func (n *nowProvider) now(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
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
	call, err := parser.ParseOperation([]byte(`{__schema__{root_fields{name}}}`))
	if err != nil {
		fmt.Println(err)
	}
	executor := executor.New(s)
	result, err := executor.HandleOperation(context.Background(), call)
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
	//     "name": "__schema__"
	//    },
	//    {
	//     "name": "__type"
	//    },
	//    {
	//     "name": "now"
	//    }
	//   ]
	//  }
	// ]
}
