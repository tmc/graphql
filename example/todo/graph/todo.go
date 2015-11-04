package graph

import (
	"github.com/tmc/graphql"
	"github.com/tmc/graphql/executor/resolver"
	"github.com/tmc/graphql/schema"
	"golang.org/x/net/context"
)

type TodoNode struct {
	Id        string
	Text      string
	Completed bool
}

type PageInfo struct {
	hasNextPage     bool
	hasPreviousPage bool
	startCursor     string
	endCursor       string
}

type TodoEdge struct {
	Node *TodoNode
}

type TodoConnection struct {
	Edges          []*TodoEdge
	TotalCount     int
	CompletedCount int
}

func (todos *TodoConnection) findIndex(todo *TodoEdge) int {
	for i, t := range todos.Edges {
		if t == todo {
			return i
		}
	}
	return -1
}

func (todos *TodoConnection) addTodo(todo *TodoEdge) {
	todos.Edges = append(todos.Edges, todo)
	todos.TotalCount += 1
	if todo.Node.Completed {
		todos.CompletedCount += 1
	}
}

func (todos *TodoConnection) removeTodo(todo *TodoEdge) {
	i := todos.findIndex(todo)
	todos.Edges = append(todos.Edges[:i], todos.Edges[i+1:]...)
	todos.TotalCount -= 1
	if todo.Node.Completed {
		todos.CompletedCount -= 1
	}
}

func (todo *TodoConnection) GraphQLTypeInfo() schema.GraphQLTypeInfo {
	return schema.GraphQLTypeInfo{
		Name:        "Todo",
		Description: "A  To Do",
		Fields: schema.GraphQLFieldSpecMap{
			"completedCount": {
				Name:        "completedCount",
				Description: "Is the ToDo completed",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					return r.Resolve(ctx, todo.CompletedCount, f)
				},
			},
			"totalCount": {
				Name:        "totalCount",
				Description: "Is the ToDo completed",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					return r.Resolve(ctx, todo.TotalCount, f)
				},
			},
			"edges": {
				Name:        "edges",
				Description: "Is the ToDo completed",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					return r.Resolve(ctx, todo.Edges, f)
				},
			},
			"pageInfo": {
				Name:        "pageInfo",
				Description: "Is the ToDo completed",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					return r.Resolve(ctx, &PageInfo{}, f)
				},
			},
		},
	}
}

func (todoedge *TodoEdge) GraphQLTypeInfo() schema.GraphQLTypeInfo {
	return schema.GraphQLTypeInfo{
		Name:        "PageInfo",
		Description: "A To Do",
		Fields: schema.GraphQLFieldSpecMap{
			"node": {
				Name:        "node",
				Description: "The id of todo.",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					return r.Resolve(ctx, todoedge.Node, f)
				},
			},
			"cursor": {
				Name:        "cursor",
				Description: "The id of todo.",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					return r.Resolve(ctx, todoedge.Node.Id, f)
				},
			},
		},
	}
}

func (pageinfo *PageInfo) GraphQLTypeInfo() schema.GraphQLTypeInfo {
	return schema.GraphQLTypeInfo{
		Name:        "PageInfo",
		Description: "A To Do",
		Fields: schema.GraphQLFieldSpecMap{
			"hasNextPage": {
				Name:        "hasNextPage",
				Description: "The id of todo.",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					return r.Resolve(ctx, pageinfo.hasNextPage, f)
				},
			},
			"hasPreviousPage": {
				Name:        "hasPreviousPage",
				Description: "The id of todo.",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					return r.Resolve(ctx, pageinfo.hasPreviousPage, f)
				},
			},
		},
	}
}

func (todonode *TodoNode) GraphQLTypeInfo() schema.GraphQLTypeInfo {
	return schema.GraphQLTypeInfo{
		Name:        "Todo",
		Description: "A  To Do",
		Fields: schema.GraphQLFieldSpecMap{
			"id": {
				Name:        "id",
				Description: "The id of todo.",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					return r.Resolve(ctx, todonode.Id, f)
				},
			},
			"text": {
				Name:        "text",
				Description: "The todo.",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					return r.Resolve(ctx, todonode.Text, f)
				},
			},
			"complete": {
				Name:        "complete",
				Description: "Is the ToDo completed",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					return r.Resolve(ctx, todonode.Completed, f)
				},
			},
		},
	}
}
