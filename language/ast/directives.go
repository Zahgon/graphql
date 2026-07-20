package ast

type Directive struct {
	Kind      string
	Loc       *Location
	Name      *Name
	Arguments []*Argument
}

func NewDirective(dir *Directive) *Directive { _ = "STUB: not implemented"; return nil }

func (dir *Directive) GetKind() string { _ = "STUB: not implemented"; return "" }

func (dir *Directive) GetLoc() *Location { _ = "STUB: not implemented"; return nil }
