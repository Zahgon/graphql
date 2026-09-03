package gqlerrors

import (
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/location"
	"github.com/graphql-go/graphql/language/source"
)

type Error struct {
	Message       string
	Stack         string
	Nodes         []ast.Node
	Source        *source.Source
	Positions     []int
	Locations     []location.SourceLocation
	OriginalError error
	Path          []interface{}
}

func (g Error) Error() string { _ = "STUB: not implemented"; return "" }

func NewError(message string, nodes []ast.Node, stack string, source *source.Source, positions []int, origError error) *Error {
	_ = "STUB: not implemented"
	return nil
}

func NewErrorWithPath(message string, nodes []ast.Node, stack string, source *source.Source, positions []int, path []interface{}, origError error) *Error {
	_ = "STUB: not implemented"
	return nil
}

func newError(message string, nodes []ast.Node, stack string, source *source.Source, positions []int, path []interface{}, origError error) *Error {
	_ = "STUB: not implemented"
	return nil
}
