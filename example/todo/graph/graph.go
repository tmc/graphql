package graph

import (
	"fmt"
	"github.com/tmc/graphql"
	"github.com/tmc/graphql/executor/resolver"
	"github.com/tmc/graphql/schema"
	"golang.org/x/net/context"
)

type Graph struct {
	nextId int
	Todos  map[int]*TodoEdge
	Users  map[string]*User
}

type AddToDoMutation struct {
	graph *Graph
	input map[string]interface{}
	edge  *TodoEdge
}

type ChangeToDoStatusMutation struct {
	graph *Graph
	input map[string]interface{}
	edge  *TodoEdge
}

type RemoveToDoMutation struct {
	graph *Graph
	input map[string]interface{}
	edge  *TodoEdge
}

type RemoveCompletedToDosMutation struct {
	graph      *Graph
	input      map[string]interface{}
	deletedIds []string
}

func NewGraph() *Graph {
	graph := &Graph{
		0,
		make(map[int]*TodoEdge),
		make(map[string]*User),
	}

	graph.Users["me"] = &User{
		"me",
		new(TodoConnection),
		new(TodoConnection),
		new(TodoConnection),
	}

	graph.AddToDo(graph.Users["me"], "Taste Javascript", false)
	graph.AddToDo(graph.Users["me"], "Buy a unicorn", false)

	return graph
}

func (graph *Graph) AddToDo(user *User, text string, complete bool) *TodoEdge {

	todo := &TodoEdge{
		&TodoNode{
			fmt.Sprintf("%d", graph.nextId),
			text,
			complete,
		},
	}
	graph.Todos[graph.nextId] = todo
	graph.nextId += 1
	user.addToDo(todo)
	return todo
}

func (graph *Graph) ChangeStatus(user *User, id string, complete bool) *TodoEdge {
	todo := user.changeStatus(id, complete)
	return todo
}

func (graph *Graph) Remove(user *User, id string) *TodoEdge {
	todo := user.remove(id)
	return todo
}

func (graph *Graph) ClearCompleted(user *User) []string {
	return user.clearCompleted()
}

func (graph *Graph) GraphQLTypeInfo() schema.GraphQLTypeInfo {
	return schema.GraphQLTypeInfo{
		Name:        "To Dos",
		Description: "A ToDo list App",
		Fields: schema.GraphQLFieldSpecMap{
			"viewer": {
				Name:        "viewer",
				Description: "A To Do user",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					g := graph.Users["me"]

					if g != nil {
						return r.Resolve(ctx, g, f)
					}
					return nil, fmt.Errorf("User not found")
				},
				IsRoot: true,
			},
			"addTodo": {
				Name:        "addToDo",
				Description: "A To Do user",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					input := ctx.Value("variables").(map[string]interface{})[f.Arguments[0].Value.(*graphql.Variable).Name].(map[string]interface{})

					todo := graph.AddToDo(graph.Users["me"], input["text"].(string), false)
					return r.Resolve(ctx, &AddToDoMutation{graph, input, todo}, f)
				},
				IsRoot: true,
			},
			"changeTodoStatus": {
				Name:        "changeTodoStatus",
				Description: "Change Todo Status",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					input := ctx.Value("variables").(map[string]interface{})[f.Arguments[0].Value.(*graphql.Variable).Name].(map[string]interface{})

					todo := graph.ChangeStatus(graph.Users["me"], input["id"].(string), input["complete"].(bool))
					return r.Resolve(ctx, &ChangeToDoStatusMutation{graph, input, todo}, f)
				},
				IsRoot: true,
			},
			"removeTodo": {
				Name:        "removeTodo",
				Description: "Change Todo Status",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					input := ctx.Value("variables").(map[string]interface{})[f.Arguments[0].Value.(*graphql.Variable).Name].(map[string]interface{})

					todo := graph.Remove(graph.Users["me"], input["id"].(string))
					return r.Resolve(ctx, &RemoveToDoMutation{graph, input, todo}, f)
				},
				IsRoot: true,
			},
			"removeCompletedTodos": {
				Name:        "removeCompletedTodos",
				Description: "Change Todo Status",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					input := ctx.Value("variables").(map[string]interface{})[f.Arguments[0].Value.(*graphql.Variable).Name].(map[string]interface{})
					deletedIds := graph.ClearCompleted(graph.Users["me"])
					return r.Resolve(ctx, &RemoveCompletedToDosMutation{graph, input, deletedIds}, f)
				},
				IsRoot: true,
			},
		},
	}
}

func (addToDo *AddToDoMutation) GraphQLTypeInfo() schema.GraphQLTypeInfo {
	return schema.GraphQLTypeInfo{
		Name:        "To Dos",
		Description: "A ToDo list App",
		Fields: schema.GraphQLFieldSpecMap{
			"clientMutationId": {
				Name:        "clientMutationId",
				Description: "A To Do user",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					input := addToDo.input
					return r.Resolve(ctx, input["clientMutationId"], f)
				},
				IsRoot: true,
			},
			"todoEdge": {
				Name:        "todoEdge",
				Description: "A To Do user",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					g := addToDo.edge

					if g != nil {
						return r.Resolve(ctx, g, f)
					}

					return nil, fmt.Errorf("Todo not found")
				},
				IsRoot: true,
			},
			"viewer": {
				Name:        "viewer",
				Description: "A To Do user",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					g := addToDo.graph.Users["me"]

					if g != nil {
						return r.Resolve(ctx, g, f)
					}
					return nil, fmt.Errorf("User not found")
				},
				IsRoot: true,
			},
		},
	}
}

func (changeTodo *ChangeToDoStatusMutation) GraphQLTypeInfo() schema.GraphQLTypeInfo {
	return schema.GraphQLTypeInfo{
		Name:        "To Dos",
		Description: "A ToDo list App",
		Fields: schema.GraphQLFieldSpecMap{
			"clientMutationId": {
				Name:        "clientMutationId",
				Description: "A To Do user",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					input := changeTodo.input
					return r.Resolve(ctx, input["clientMutationId"], f)
				},
				IsRoot: true,
			},
			"todo": {
				Name:        "todo",
				Description: "A To Do user",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					g := changeTodo.edge.Node

					if g != nil {
						return r.Resolve(ctx, g, f)
					}

					return nil, fmt.Errorf("Todo not found")
				},
				IsRoot: true,
			},
			"viewer": {
				Name:        "viewer",
				Description: "A To Do user",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					g := changeTodo.graph.Users["me"]

					if g != nil {
						return r.Resolve(ctx, g, f)
					}
					return nil, fmt.Errorf("User not found")
				},
				IsRoot: true,
			},
		},
	}
}

func (removeTodo *RemoveToDoMutation) GraphQLTypeInfo() schema.GraphQLTypeInfo {
	return schema.GraphQLTypeInfo{
		Name:        "To Dos",
		Description: "A ToDo list App",
		Fields: schema.GraphQLFieldSpecMap{
			"clientMutationId": {
				Name:        "clientMutationId",
				Description: "A To Do user",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					input := removeTodo.input
					return r.Resolve(ctx, input["clientMutationId"], f)
				},
				IsRoot: true,
			},
			"deletedTodoId": {
				Name:        "deletedTodoId",
				Description: "A To Do user",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					g := removeTodo.edge.Node.Id

					return r.Resolve(ctx, g, f)
				},
				IsRoot: true,
			},
			"viewer": {
				Name:        "viewer",
				Description: "A To Do user",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					g := removeTodo.graph.Users["me"]

					if g != nil {
						return r.Resolve(ctx, g, f)
					}
					return nil, fmt.Errorf("User not found")
				},
				IsRoot: true,
			},
		},
	}
}

func (removeCompletedTodo *RemoveCompletedToDosMutation) GraphQLTypeInfo() schema.GraphQLTypeInfo {
	return schema.GraphQLTypeInfo{
		Name:        "To Dos",
		Description: "A ToDo list App",
		Fields: schema.GraphQLFieldSpecMap{
			"clientMutationId": {
				Name:        "clientMutationId",
				Description: "A To Do user",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					input := removeCompletedTodo.input
					return r.Resolve(ctx, input["clientMutationId"], f)
				},
				IsRoot: true,
			},
			"deletedTodoIds": {
				Name:        "deletedTodoId",
				Description: "A To Do user",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					g := removeCompletedTodo.deletedIds

					return r.Resolve(ctx, g, f)
				},
				IsRoot: true,
			},
			"viewer": {
				Name:        "viewer",
				Description: "A To Do user",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					g := removeCompletedTodo.graph.Users["me"]

					if g != nil {
						return r.Resolve(ctx, g, f)
					}
					return nil, fmt.Errorf("User not found")
				},
				IsRoot: true,
			},
		},
	}
}
