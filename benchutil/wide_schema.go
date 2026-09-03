package benchutil

import (
	"github.com/graphql-go/graphql"
)

func WideSchemaWithXFieldsAndYItems(x int, y int) graphql.Schema {
	_ = "STUB: not implemented"
	return *new(graphql.Schema)
}

func generateXWideFields(x int) graphql.Fields {
	_ = "STUB: not implemented"
	return *new(graphql.Fields)
}

func generateWideFieldFromX(x int) *graphql.Field { _ = "STUB: not implemented"; return nil }

func generateWideTypeFromX(x int) graphql.Type {
	_ = "STUB: not implemented"
	return *new(graphql.Type)
}

func generateFieldNameFromX(x int) string { _ = "STUB: not implemented"; return "" }

func generateWideResolveFromX(x int) func(p graphql.ResolveParams) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil
}

func WideSchemaQuery(x int) string { _ = "STUB: not implemented"; return "" }

func WideArgedSchemaWithXFieldsAndYItems(x int, y int) graphql.Schema {
	_ = "STUB: not implemented"
	return *new(graphql.Schema)
}

func generateXArgedWideFields(x int) graphql.Fields {
	_ = "STUB: not implemented"
	return *new(graphql.Fields)
}

func generateArgedWideFieldFromX(x int) *graphql.Field { _ = "STUB: not implemented"; return nil }

func generateArgedWideResolveFromX(x int) func(p graphql.ResolveParams) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil
}

func WideArgedSchemaQueryWithVariable(x int) string { _ = "STUB: not implemented"; return "" }

func WideArgedSchemaQueryWithLiteral(x int, literal string) string {
	_ = "STUB: not implemented"
	return ""
}
