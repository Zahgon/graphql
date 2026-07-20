package ast

type Document struct {
	Kind        string
	Loc         *Location
	Definitions []Node
}

func NewDocument(d *Document) *Document { _ = "STUB: not implemented"; return nil }

func (node *Document) GetKind() string { _ = "STUB: not implemented"; return "" }

func (node *Document) GetLoc() *Location { _ = "STUB: not implemented"; return nil }
