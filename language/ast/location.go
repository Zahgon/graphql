package ast

import (
	"github.com/graphql-go/graphql/language/source"
)

type Location struct {
	Start  int
	End    int
	Source *source.Source
}

func NewLocation(loc *Location) *Location { _ = "STUB: not implemented"; return nil }
