package schema

import (
	"github.com/graphql-go/graphql"
)

var TodoList []Todo

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string { _ = "STUB: not implemented"; return "" }

var todoType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Todo",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"text": &graphql.Field{
			Type: graphql.String,
		},
		"done": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{

		"createTodo": &graphql.Field{
			Type:        todoType,
			Description: "Create new todo",
			Args: graphql.FieldConfigArgument{
				"text": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {

				text, _ := params.Args["text"].(string)

				newID := RandStringRunes(8)

				newTodo := Todo{
					ID:   newID,
					Text: text,
					Done: false,
				}

				TodoList = append(TodoList, newTodo)

				return newTodo, nil
			},
		},

		"updateTodo": &graphql.Field{
			Type:        todoType,
			Description: "Update existing todo, mark it done or not done",
			Args: graphql.FieldConfigArgument{
				"done": &graphql.ArgumentConfig{
					Type: graphql.Boolean,
				},
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {

				done, _ := params.Args["done"].(bool)
				id, _ := params.Args["id"].(string)
				affectedTodo := Todo{}

				for i := 0; i < len(TodoList); i++ {
					if TodoList[i].ID == id {
						TodoList[i].Done = done

						affectedTodo = TodoList[i]
						break
					}
				}

				return affectedTodo, nil
			},
		},
	},
})

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{

		"todo": &graphql.Field{
			Type:        todoType,
			Description: "Get single todo",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {

				idQuery, isOK := params.Args["id"].(string)
				if isOK {

					for _, todo := range TodoList {
						if todo.ID == idQuery {
							return todo, nil
						}
					}
				}

				return Todo{}, nil
			},
		},

		"lastTodo": &graphql.Field{
			Type:        todoType,
			Description: "Last todo added",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return TodoList[len(TodoList)-1], nil
			},
		},

		"todoList": &graphql.Field{
			Type:        graphql.NewList(todoType),
			Description: "List of todos",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return TodoList, nil
			},
		},
	},
})

var TodoSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})
