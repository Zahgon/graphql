package gqlerrors

import (
	"github.com/graphql-go/graphql/language/location"
	"github.com/graphql-go/graphql/language/source"
)

func NewSyntaxError(s *source.Source, position int, description string) *Error {
	_ = "STUB: not implemented"
	return nil
}

func printCharCode(code rune) string { _ = "STUB: not implemented"; return "" }

func printLine(str string) string { _ = "STUB: not implemented"; return "" }

func highlightSourceAtLocation(s *source.Source, l location.SourceLocation) string {
	_ = "STUB: not implemented"
	return ""
}

func lpad(l int, s string) string { _ = "STUB: not implemented"; return "" }
