package executor

import (
	"fmt"
	"log"
	"reflect"
	"sync"

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

type fieldResult struct {
	FieldName string
	Value     interface{}
	Err       error
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
	results := make(chan fieldResult)
	wg := sync.WaitGroup{}

	for _, selection := range field.SelectionSet {
		fieldName := selection.Field.Name
		fieldHandler, ok := typeInfo.Fields[fieldName]
		if !ok {
			return nil, fmt.Errorf("No handler for field '%s' on type '%T'", fieldName, graphQLValue)
		}
		wg.Add(1)
		go func(selection graphql.Selection) {
			defer wg.Done()
			partial, err := fieldHandler.Func(ctx, e, selection.Field)
			if err != nil {
				results <- fieldResult{Err: err}
				return
			}
			resolved, err := e.Resolve(ctx, partial, selection.Field)
			if err != nil {
				results <- fieldResult{Err: err}
				return
			}
			if selection.Field.Alias != "" {
				fieldName = selection.Field.Alias
			}
			results <- fieldResult{
				FieldName: fieldName, Value: resolved, Err: err,
			}
		}(selection)
	}
	go func() {
		wg.Wait()
		close(results)
	}()
	for r := range results {
		if r.Err != nil {
			return nil, r.Err
		}
		result[r.FieldName] = r.Value
	}
	return result, nil
}

func (e *Executor) resolveSlice(ctx context.Context, partials interface{}, field *graphql.Field) (interface{}, error) {
	v := reflect.ValueOf(partials)
	results := make([]interface{}, 0, v.Len())
	resChan := make(chan fieldResult)
	wg := sync.WaitGroup{}
	for i := 0; i < v.Len(); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			result, err := e.Resolve(ctx, v.Index(i).Interface(), field)
			resChan <- fieldResult{Value: result, Err: err}
		}(i)
	}
	go func() {
		wg.Wait()
		close(resChan)
	}()
	for result := range resChan {
		if result.Err != nil {
			return nil, result.Err
		}
		results = append(results, result.Value)
	}
	return results, nil
}
