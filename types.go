package graphql

// Call represents a root call into a graphql schema
//
// Example string representation:
// node(42){fieldX}
type Call struct {
	Name      string
	Arguments Arguments `json:",omitempty"`
	Fields    Fields    `json:",omitempty"`
}

// Argument is an argument to a Call
type Argument string

// Arguments is a collection of Argument values
type Arguments []Argument

// Fields is a collection of Field values
type Fields []Field

// Field represents a named field, a field call, or a set of Fields for a sub-object.
type Field struct {
	Call   *Call  `json:",omitempty"`
	Name   string `json:",omitempty"`
	Fields Fields `json:",omitempty"`
}
