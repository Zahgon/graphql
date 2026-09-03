package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
)

var Schema graphql.Schema

var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"me": &graphql.Field{
				Type: userType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return p.Context.Value("currentUser"), nil
				},
			},
		},
	})

func graphqlHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func main() {
	http.HandleFunc("/graphql", graphqlHandler)
	fmt.Println("Now server is running on port 8080")
	fmt.Println("Test with Get      : curl -g 'http://localhost:8080/graphql?query={me{id,name}}'")
	http.ListenAndServe(":8080", nil)
}

func init() {
	s, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
	if err != nil {
		log.Fatalf("failed to create schema, error: %v", err)
	}
	Schema = s
}
