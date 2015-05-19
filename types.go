package graphql

import "encoding/json"

type OperationType string

const (
	OperationQuery    OperationType = "query"
	OperationMutation OperationType = "mutation"
)

type Document struct {
	Operations          []Operation
	FragmentDefinitions []FragmentDefinition `json:",omitempty"`
	EnumDefinitions     []EnumDefinition     `json:",omitempty"`
	TypeDefinitions     []TypeDefinition     `json:",omitempty"`
	TypeExtensions      []TypeExtension      `json:",omitempty"`
}

type Operation struct {
	Type                OperationType        `json:",omitempty"`
	Name                string               `json:",omitempty"`
	Selections          []Selection          `json:",omitempty"`
	VariableDefinitions []VariableDefinition `json:",omitempty"`
	Directives          []Directive          `json:",omitempty"`
}

func (o *Operation) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}

// A Selection is either a Field or a FragmentSpread
type Selection struct {
	Field          *Field          `json:",omitempty"`
	FragmentSpread *FragmentSpread `json:",omitempty"`
}

func (s *Selection) String() string {
	j, _ := json.Marshal(s)
	return string(j)
}

type Field struct {
	Name       string      `json:",omitempty"`
	Arguments  Arguments   `json:",omitempty"`
	Selections Selections  `json:",omitempty"`
	Alias      string      `json:",omitempty"`
	Directives []Directive `json:",omitempty"`
}

type FragmentSpread struct {
	Name       string      `json:",omitempty"`
	Directives []Directive `json:",omitempty"`
}

// Argument is an argument to a Field Call
type Argument struct {
	Name  string
	Value interface{}
}

// Arguments is a collection of Argument values
type Arguments []Argument

// Fields is a collection of Field values
type Selections []Selection

// Fragments

type FragmentDefinition struct {
	Name       string
	Type       Type
	Selections []Selection
	Directives []Directive `json:",omitempty"`
}

// Type system

type TypeDefinition struct {
	Name             string
	Interfaces       []Interface `json:",omitempty"`
	FieldDefinitions []FieldDefinition
}

type TypeExtension struct {
	Name             string
	Interfaces       []Interface `json:",omitempty"`
	FieldDefinitions []FieldDefinition
}

type FieldDefinition struct {
	Name                string
	Type                Type
	ArgumentDefinitions []ArgumentDefinition `json:",omitempty"`
}

type ArgumentDefinition struct {
	Name         string
	Type         Type
	DefaultValue *Value `json:",omitempty"`
}

type Type struct {
	Name     string
	Optional bool
	Params   []Type `json:",omitempty"`
}

type Value interface{}

type Interface struct{}

// Enums

type EnumDefinition struct {
	Name   string
	Values []string
}

type EnumValue struct {
	EnumTypeName string
	Value        string
}

// Variables

type VariableDefinition struct {
	Variable     Variable
	Type         Type
	DefaultValue *Value `json:",omitempty"`
}

type Variable struct {
	Name              string
	PropertySelection *Variable `json:",omitempty"`
}

// Directives

type Directive struct {
	Name  string
	Type  *Type  `json:",omitempty"`
	Value *Value `json:",omitempty"`
}
