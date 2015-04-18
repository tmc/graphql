package graphql

type Call struct {
	Name      string
	Arguments Arguments
	Fields    Fields `json:",omitempty"`
}

type Argument string

type Arguments []Argument

type Fields []Field

type Field struct {
	Name   string `json:",omitempty"`
	Fields Fields `json:",omitempty"`
}
