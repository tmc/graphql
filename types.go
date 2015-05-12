package graphql

type OperationType string

const (
	OperationQuery    OperationType = "query"
	OperationMutation OperationType = "mutation"
)

type Operation struct {
	Type       OperationType `json:",omitempty"`
	Name       string        `json:",omitempty"`
	Selections []Selection   `json:",omitempty"`
	//Variables []Variable `json:",omitempty"`
}

type Selection struct {
	FieldName  string     `json:",omitempty"`
	Arguments  Arguments  `json:",omitempty"`
	Selections Selections `json:",omitempty"`
	// FieldAlias string
	// Directives []Directive
}

// Argument is an argument to a Field Call
type Argument struct {
	Name  string
	Value string
}

// Arguments is a collection of Argument values
type Arguments []Argument

// Fields is a collection of Field values
type Selections []Selection
