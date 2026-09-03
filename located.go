package graphql

import (
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/graphql-go/graphql/language/ast"
)

func NewLocatedError(err interface{}, nodes []ast.Node) *gqlerrors.Error {
	_ = "STUB: not implemented"
	return nil
}

func NewLocatedErrorWithPath(err interface{}, nodes []ast.Node, path []interface{}) *gqlerrors.Error {
	_ = "STUB: not implemented"
	return nil
}

func newLocatedError(err interface{}, nodes []ast.Node, path []interface{}) *gqlerrors.Error {
	_ = "STUB: not implemented"
	return nil
}

func FieldASTsToNodeASTs(fieldASTs []*ast.Field) []ast.Node { _ = "STUB: not implemented"; return nil }
