package graphql

import "encoding/json"

// OperationType is either "query" or "mutation"
// Queries are reads and mutations cause side-effects.
type OperationType string

const (
	// OperationQuery is a read operation.
	OperationQuery OperationType = "query"
	// OperationMutation is a mutation.
	OperationMutation OperationType = "mutation"
)

// Document is the top-level representation of a string in GraphQL.
type Document struct {
	Operations          []Operation
	FragmentDefinitions []FragmentDefinition `json:",omitempty"`
	EnumDefinitions     []EnumDefinition     `json:",omitempty"`
	TypeDefinitions     []TypeDefinition     `json:",omitempty"`
	TypeExtensions      []TypeExtension      `json:",omitempty"`
}

// Operation is either a read or mutation in GraphQL.
type Operation struct {
	Type                OperationType        `json:",omitempty"`
	Name                string               `json:",omitempty"`
	SelectionSet        SelectionSet         `json:",omitempty"`
	VariableDefinitions []VariableDefinition `json:",omitempty"`
	Directives          []Directive          `json:",omitempty"`
}

func (o *Operation) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

// A Selection is either a Field, a FragmentSpread, or an InlineFragment
type Selection struct {
	Field          *Field          `json:",omitempty"`
	FragmentSpread *FragmentSpread `json:",omitempty"`
	InlineFragment *InlineFragment `json:",omitempty"`
}

func (s *Selection) String() string {
	j, _ := json.Marshal(s)
	return string(j)
}

// A Field is one of the most important concepts in GraphQL. Fields specify what
// parts of data you would like to select.
type Field struct {
	Name         string       `json:",omitempty"`
	Arguments    Arguments    `json:",omitempty"`
	SelectionSet SelectionSet `json:",omitempty"`
	Alias        string       `json:",omitempty"`
	Directives   []Directive  `json:",omitempty"`
}

// FragmentSpread is a reference to a QueryFragment elsewhere in an Operation.
type FragmentSpread struct {
	Name       string      `json:",omitempty"`
	Directives []Directive `json:",omitempty"`
}

// InlineFragment is used in-line to apply a type condition within a selection.
type InlineFragment struct {
	TypeCondition string      `json:",omitempty"`
	Directives    []Directive `json:",omitempty"`
	SelectionSet  SelectionSet
}

// Argument is an argument to a Field Call.
type Argument struct {
	Name  string
	Value interface{}
}

// Arguments is a collection of Argument values
type Arguments []Argument

// Get is a helper to fetch a particular argument by name.
func (a Arguments) Get(name string) (interface{}, bool) {
	for _, arg := range a {
		if arg.Name == name {
			return arg.Value, true
		}
	}
	return nil, false
}

// SelectionSet is a collection of Selection
type SelectionSet []Selection

// Fragments

// FragmentDefinition defines a Query Fragment
type FragmentDefinition struct {
	Name          string
	TypeCondition string
	SelectionSet  SelectionSet
	Directives    []Directive `json:",omitempty"`
}

// Type system

// TypeDefinition defines a type.
type TypeDefinition struct {
	Name             string
	Interfaces       []Interface `json:",omitempty"`
	FieldDefinitions []FieldDefinition
}

// TypeExtension extends an existing type.
type TypeExtension struct {
	Name             string
	Interfaces       []Interface `json:",omitempty"`
	FieldDefinitions []FieldDefinition
}

// FieldDefinition defines a fields on a type.
type FieldDefinition struct {
	Name                string
	Type                Type
	ArgumentDefinitions []ArgumentDefinition `json:",omitempty"`
}

// ArgumentDefinition defines an argument for a field on a type.
type ArgumentDefinition struct {
	Name         string
	Type         Type
	DefaultValue *Value `json:",omitempty"`
}

// Type describes an argument's type.
type Type struct {
	Name     string
	Optional bool
	Params   []Type `json:",omitempty"`
}

// Value refers to a value
type Value interface{}

// Interface descibes a set of methods a type must conform to to satisfy it.
// TODO
type Interface struct{}

// Enums

// EnumDefinition defines an enum.
type EnumDefinition struct {
	Name   string
	Values []string
}

// EnumValue describes a possible value for an enum.
type EnumValue struct {
	EnumTypeName string
	Value        string
}

// Variables

// VariableDefinition defines a variable for an Operation.
type VariableDefinition struct {
	Variable     Variable
	Type         Type
	DefaultValue *Value `json:",omitempty"`
}

// Variable describes a reference to a variable.
type Variable struct {
	Name              string
	PropertySelection *Variable `json:",omitempty"`
}

// Directives

// Directive describes a directive which can alter behavior in different parts of a GraphQL Operation.
type Directive struct {
	Name  string
	Type  *Type  `json:",omitempty"`
	Value *Value `json:",omitempty"`
}
