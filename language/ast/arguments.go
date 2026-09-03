package ast

type Argument struct {
	Kind  string
	Loc   *Location
	Name  *Name
	Value Value
}

func NewArgument(arg *Argument) *Argument { _ = "STUB: not implemented"; return nil }

func (arg *Argument) GetKind() string { _ = "STUB: not implemented"; return "" }

func (arg *Argument) GetLoc() *Location { _ = "STUB: not implemented"; return nil }
