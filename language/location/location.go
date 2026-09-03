package location

import (
	"github.com/graphql-go/graphql/language/source"
)

type SourceLocation struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}

func GetLocation(s *source.Source, position int) SourceLocation {
	_ = "STUB: not implemented"
	return *new(SourceLocation)
}
