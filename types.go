package graphql

// Call represents a root call into a graphql schema
//
// Example string representation:
// node(42){fieldX}
type Call struct {
	Name      string
	Arguments Arguments
	Fields    Fields `json:",omitempty"`
}

// Argument is an argument to a Call
type Argument string

// Arguments is a collection of Argument values
type Arguments []Argument

// Fields is a collection of Field values
type Fields []Field

// Field represents either a named field or a set of Fields for a sub-object.
type Field struct {
	Name   string `json:",omitempty"`
	Fields Fields `json:",omitempty"`
}
