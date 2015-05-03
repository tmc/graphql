package parser_test

import (
	"testing"

	"github.com/tmc/graphql/internal/parser"
)

var shouldParse = []string{
	`foo()`,
	`{foo()}`,
	`node(id:42)`,
	`foo(){id}`,
	`foo(id1:1,id2:2){id}`,
	`foo(id2:1,id2:2){id,{nest,some,{fields}}}`,
	`node(id:42){id,friends.top(10)}`,
	`node(id:42) { id , friends.top(10) }`,
	`{
	  user(id: 3500401) {
	    id,
	    name,
	    isViewerFriend,
	    profilePicture(size: 50)  {
	      uri,
	      width,
	      height
	    }
	  }
	}
	`,
}

func TestSuccessfulParses(t *testing.T) {
	for i, in := range shouldParse {
		_, err := parser.Parse("parser_test.go", []byte(in))
		if err != nil {
			t.Errorf("case %d: %v", i+1, err)
		}
	}
}
