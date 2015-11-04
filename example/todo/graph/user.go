package graph

import (
	"github.com/tmc/graphql"
	"github.com/tmc/graphql/executor/resolver"
	"github.com/tmc/graphql/schema"
	"golang.org/x/net/context"
)

type User struct {
	Id             string
	AnyTodos       *TodoConnection
	CompletedTodos *TodoConnection
	ActiveTodos    *TodoConnection
}

func (user *User) addToDo(todo *TodoEdge) {
	user.AnyTodos.addTodo(todo)
	if todo.Node.Completed {
		user.CompletedTodos.addTodo(todo)
	} else {
		user.ActiveTodos.addTodo(todo)
	}
}

func (user *User) changeStatus(id string, complete bool) *TodoEdge {
	var todo *TodoEdge
	for _, t := range user.AnyTodos.Edges {
		if t.Node.Id == id {
			todo = t
		}
	}
	if todo.Node.Completed != complete {
		todo.Node.Completed = complete
		if complete {
			user.CompletedTodos.addTodo(todo)
			user.ActiveTodos.removeTodo(todo)
			user.AnyTodos.CompletedCount += 1
		} else {
			user.ActiveTodos.addTodo(todo)
			user.CompletedTodos.removeTodo(todo)
			user.AnyTodos.CompletedCount -= 1
		}
	}
	return todo
}

func (user *User) remove(id string) *TodoEdge {
	var todo *TodoEdge
	for _, t := range user.AnyTodos.Edges {
		if t.Node.Id == id {
			todo = t
		}
	}
	user.AnyTodos.removeTodo(todo)
	if todo.Node.Completed {
		user.CompletedTodos.removeTodo(todo)
	} else {
		user.ActiveTodos.removeTodo(todo)
	}
	return todo
}

func (user *User) clearCompleted() []string {
	removedIds := make([]string, len(user.CompletedTodos.Edges), len(user.CompletedTodos.Edges))
	for i, t := range user.CompletedTodos.Edges {
		removedIds[i] = t.Node.Id
		user.AnyTodos.removeTodo(t)
		user.CompletedTodos.removeTodo(t)
	}
	return removedIds
}

func (user *User) GraphQLTypeInfo() schema.GraphQLTypeInfo {
	return schema.GraphQLTypeInfo{
		Name:        "User",
		Description: "A user",
		Fields: schema.GraphQLFieldSpecMap{
			"id": {
				Name:        "id",
				Description: "The id of user.",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					return r.Resolve(ctx, user.Id, f)
				},
			},
			"todos": {
				Name:        "todos",
				Description: "The todos for a user.",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {

					if status, ok := f.Arguments.Get("status"); ok {
						if status == "completed" {
							return r.Resolve(ctx, user.CompletedTodos, f)
						}
						if status == "active" {
							return r.Resolve(ctx, user.ActiveTodos, f)
						}
					}

					// if f.Arguments[0].Name == "status" {
					// 	if f.Arguments[0].Value.(graphql.Value).Value == "completed" {
					// 		return r.Resolve(ctx, user.CompletedTodos, f)
					// 	}
					// 	if f.Arguments[0].Value.(*graphql.Value).Value == "active" {
					// 		return r.Resolve(ctx, user.ActiveTodos, f)
					// 	}
					// }
					return r.Resolve(ctx, user.AnyTodos, f)
				},
			},
			"completedCount": {
				Name:        "completedCount",
				Description: "The todos for a user.",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					return r.Resolve(ctx, user.AnyTodos.CompletedCount, f)
				},
			},
			"totalCount": {
				Name:        "totalCount",
				Description: "The todos for a user.",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					return r.Resolve(ctx, user.AnyTodos.TotalCount, f)
				},
			},
		},
	}
}
