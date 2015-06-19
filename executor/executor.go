package executor

import (
	"fmt"
	"log"
	"reflect"

	"github.com/tmc/graphql"
	"github.com/tmc/graphql/schema"
	"golang.org/x/net/context"
)

type Executor struct {
	schema *schema.Schema
}

func New(schema *schema.Schema) *Executor {
	return &Executor{
		schema: schema,
	}
}

func (e *Executor) HandleOperation(ctx context.Context, o *graphql.Operation) (interface{}, error) {
	rootSelections := o.SelectionSet
	rootFields := e.schema.RootFields()
	result := make([]interface{}, 0)

	for _, selection := range rootSelections {
		rootFieldHandler, ok := rootFields[selection.Field.Name]
		if !ok {
			return nil, fmt.Errorf("Root field '%s' is not registered", selection.Field.Name)
		}
		partial, err := rootFieldHandler.Func(ctx, e, selection.Field)
		if err != nil {
			return nil, err
		}
		resolved, err := e.Resolve(ctx, partial, selection.Field)
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

func (e *Executor) Resolve(ctx context.Context, partial interface{}, field *graphql.Field) (interface{}, error) {
	if partial != nil && isSlice(partial) {
		return e.resolveSlice(ctx, partial, field)
	}
	graphQLValue, ok := partial.(schema.GraphQLType)
	// if we have a scalar we're done
	if !ok {
		log.Printf("returning scalar %T: %v\n", partial, partial)
		return partial, nil
	}
	// check against returning object as non-leaf
	if len(field.SelectionSet) == 0 {
		return nil, fmt.Errorf("Cannot return a '%T' as a leaf", graphQLValue)
	}

	result := map[string]interface{}{}
	typeInfo := schema.WithIntrospectionField(graphQLValue.GraphQLTypeInfo())
	for _, selection := range field.SelectionSet {
		fieldName := selection.Field.Name
		fieldHandler, ok := typeInfo.Fields[fieldName]
		if !ok {
			return nil, fmt.Errorf("No handler for field '%s' on type '%T'", fieldName, graphQLValue)
		}
		partial, err := fieldHandler.Func(ctx, e, selection.Field)
		if err != nil {
			return nil, err // TODO(tmc) decorate error
		}
		resolved, err := e.Resolve(ctx, partial, selection.Field)
		if err != nil {
			return nil, err
		}
		if selection.Field.Alias != "" {
			fieldName = selection.Field.Alias
		}
		result[fieldName] = resolved
	}
	return result, nil
}

func (e *Executor) resolveSlice(ctx context.Context, partials interface{}, field *graphql.Field) (interface{}, error) {
	v := reflect.ValueOf(partials)
	results := make([]interface{}, 0, v.Len())
	for i := 0; i < v.Len(); i++ {
		result, err := e.Resolve(ctx, v.Index(i).Interface(), field)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}
