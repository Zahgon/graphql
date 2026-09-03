package ast

type Name struct {
	Kind  string
	Loc   *Location
	Value string
}

func NewName(node *Name) *Name { _ = "STUB: not implemented"; return nil }

func (node *Name) GetKind() string { _ = "STUB: not implemented"; return "" }

func (node *Name) GetLoc() *Location { _ = "STUB: not implemented"; return nil }
