package gqlerrors

import (
	"github.com/graphql-go/graphql/language/ast"
)

func NewLocatedError(err interface{}, nodes []ast.Node) *Error {
	_ = "STUB: not implemented"
	return nil
}

func FieldASTsToNodeASTs(fieldASTs []*ast.Field) []ast.Node { _ = "STUB: not implemented"; return nil }
