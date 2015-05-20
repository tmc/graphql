package executor

import (
	"fmt"
	"log"
	"reflect"

	"github.com/tmc/graphql"
	"github.com/tmc/graphql/schema"
)

type Executor struct {
	schema *schema.Schema
}

func New(schema *schema.Schema) *Executor {
	return &Executor{
		schema: schema,
	}
}

func (e *Executor) HandleOperation(o *graphql.Operation) (interface{}, error) {
	rootSelections := o.Selections
	rootCalls := e.schema.RootCalls()
	result := make([]interface{}, 0)

	for _, selection := range rootSelections {
		rootCallHandler, ok := rootCalls[selection.Field.Name]
		if !ok {
			return nil, fmt.Errorf("Root call '%s' is not registered", selection.Field.Name)
		}
		partial, err := rootCallHandler.Func(e, selection.Field)
		if err != nil {
			return nil, err
		}
		resolved, err := e.Resolve(partial, selection.Field)
		if err != nil {
			return nil, err
		}
		result = append(result, resolved)
	}
	return result, nil
}

func isSlice(value interface{}) bool {
	if value == nil {
		return false
	}
	return reflect.TypeOf(value).Kind() == reflect.Slice
}

func (e *Executor) Resolve(partial interface{}, field *graphql.Field) (interface{}, error) {
	if partial != nil && isSlice(partial) {
		return e.resolveSlice(partial, field)
	}
	graphQLValue, ok := partial.(schema.GraphQLType)
	// if we have a scalar we're done
	if !ok {
		log.Printf("returning scalar %T: %v\n", partial, partial)
		return partial, nil
	}
	// check against returning object as non-leaf
	if len(field.Selections) == 0 {
		return nil, fmt.Errorf("Cannot return a '%T' as a leaf", graphQLValue)
	}

	result := map[string]interface{}{}
	typeInfo := schema.WithIntrospectionField(graphQLValue.GraphQLTypeInfo())
	for _, selection := range field.Selections {
		fieldName := selection.Field.Name
		fieldHandler, ok := typeInfo.Fields[fieldName]
		if !ok {
			return nil, fmt.Errorf("No handler for field '%s' on type '%T'", fieldName, graphQLValue)
		}
		partial, err := fieldHandler.Func(e, selection.Field)
		if err != nil {
			return nil, err // TODO(tmc) decorate error
		}
		resolved, err := e.Resolve(partial, selection.Field)
		if err != nil {
			return nil, err
		}
		result[fieldName] = resolved
	}
	return result, nil
}

func (e *Executor) resolveSlice(partials interface{}, field *graphql.Field) (interface{}, error) {
	v := reflect.ValueOf(partials)
	results := make([]interface{}, 0, v.Len())
	for i := 0; i < v.Len(); i++ {
		result, err := e.Resolve(v.Index(i).Interface(), field)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}
