package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

type NullString struct {
	sql.NullString
}

func (v NullString) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (v *NullString) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewNullString(value string) *NullString { _ = "STUB: not implemented"; return nil }

func SerializeNullString(value interface{}) interface{} { _ = "STUB: not implemented"; return nil }

func ParseNullString(value interface{}) interface{} { _ = "STUB: not implemented"; return nil }

func ParseLiteralNullString(valueAST ast.Value) interface{} { _ = "STUB: not implemented"; return nil }

var NullableString = graphql.NewScalar(graphql.ScalarConfig{
	Name:         "NullableString",
	Description:  "The `NullableString` type repesents a nullable SQL string.",
	Serialize:    SerializeNullString,
	ParseValue:   ParseNullString,
	ParseLiteral: ParseLiteralNullString,
})

type Person struct {
	Name        string      `json:"name"`
	FavoriteDog *NullString `json:"favorite_dog"`
}

var PersonType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Person",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"favorite_dog": &graphql.Field{
			Type: NullableString,
		},
	},
})

func main() {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"people": &graphql.Field{
					Type: graphql.NewList(PersonType),
					Args: graphql.FieldConfigArgument{
						"favorite_dog": &graphql.ArgumentConfig{
							Type: NullableString,
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						dog, dogOk := p.Args["favorite_dog"].(*NullString)
						people := []Person{
							Person{Name: "Alice", FavoriteDog: NewNullString("Yorkshire Terrier")},

							Person{Name: "Bob", FavoriteDog: NewNullString("")},
							Person{Name: "Chris", FavoriteDog: NewNullString("French Bulldog")},
						}
						switch {
						case dogOk:
							log.Printf("favorite_dog from arguments: %+v", dog)
							dogPeople := make([]Person, 0)
							for _, p := range people {
								if p.FavoriteDog.Valid {
									if p.FavoriteDog.String == dog.String {
										dogPeople = append(dogPeople, p)
									}
								}
							}
							return dogPeople, nil
						default:
							return people, nil
						}
					},
				},
			},
		}),
	})
	if err != nil {
		log.Fatal(err)
	}
	query := `
query {
  people {
    name
    favorite_dog
    }
}`
	queryWithArgument := `
query {
  people(favorite_dog: "Yorkshire Terrier") {
    name
    favorite_dog
  }
}`
	r1 := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	r2 := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: queryWithArgument,
	})
	if len(r1.Errors) > 0 {
		log.Fatal(r1)
	}
	if len(r2.Errors) > 0 {
		log.Fatal(r1)
	}
	b1, err := json.MarshalIndent(r1, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	b2, err := json.MarshalIndent(r2, "", "  ")
	if err != nil {
		log.Fatal(err)

	}
	fmt.Printf("\nQuery: %+v\n", string(query))
	fmt.Printf("\nResult: %+v\n", string(b1))
	fmt.Printf("\nQuery (with arguments): %+v\n", string(queryWithArgument))
	fmt.Printf("\nResult (with arguments): %+v\n", string(b2))
}
