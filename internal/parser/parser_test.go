package parser_test

import (
	"testing"

	"github.com/tmc/graphql/internal/parser"
)

var shouldParse = []string{
	`# "me" usually represents the currently logged in user.
	query getMe {
	  me
	}`,
	`query getFoobar {
	  user(id: 42) {
	      id,
	      firstName,
	      lastName
	    }
	}`,
	`{
	  user(id: 42) {
		id,
		name,
		profilePic(width: 100, height: 50)
	  }
	}`,
	`query getFoobarFriends($cursor: String) {
	  user(id: 42) {
		friends(isViewerFriend: true, first: 10, after: $cursor) {
		  nodes {
			name
		  }
		}
	  }
	}`,
	`{
	  node(username: "zuck") @expect: User {
		friends { count }
	  }
	}`,
	`query myQuery($someTest: Boolean) {
	  experimental_field @if: $someTest,
	  control_field @unless: $someTest
	}`,
	`{
	  user(id: 42) {
		friends(first: 10) {
		  ...friendFields
		},
		mutualFriends(first: 10) {
		  ...friendFields
		}
	  }
	}

	fragment User friendFields {
	  id,
	  name,
	  profilePic(size: 50)
	}`,
	`query getCommentThread($threadID: String) {
	  thread(id: $threadID) {
		...threadComments
	  }
	}

	fragment Comment ThreadComments @maxDepth: 5 {
	  comments(first: 5) {
		nodes {
		  author {
			profilePic
		  },
		  ...ThreadComments
		}
	  }
	}`,
	`extend User {
	  currentLocation: GPSCoordinate
	}

	type GPSCoordinate {
	  lat: Number
	  lon: Number
	}

	enum Color { RED, GREEN, BLUE }

	query Complement(color: Color = Color.GREEN){}

	extend User {
	  # Resolution is in meters
	  currentLocation(resolution: Int = 3000): GPSCoordinate
	}`,
	`type Person {
	  name: String
	  age: Int
	  picture: Url
	}`,
}

func TestSuccessfulParses(t *testing.T) {
	for i, in := range shouldParse {
		//_, err := parser.Parse("parser_test.go", []byte(in), parser.Debug(true))
		_, err := parser.Parse("parser_test.go", []byte(in))
		if err != nil {
			t.Errorf("case %d: %v", i+1, err)
		}
	}
}
