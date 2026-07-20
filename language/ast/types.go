package ast

type Type interface {
	GetKind() string
	GetLoc() *Location
	String() string
}

var _ Type = (*Named)(nil)
var _ Type = (*List)(nil)
var _ Type = (*NonNull)(nil)

type Named struct {
	Kind string
	Loc  *Location
	Name *Name
}

func NewNamed(t *Named) *Named { _ = "STUB: not implemented"; return nil }

func (t *Named) GetKind() string { _ = "STUB: not implemented"; return "" }

func (t *Named) GetLoc() *Location { _ = "STUB: not implemented"; return nil }

func (t *Named) String() string { _ = "STUB: not implemented"; return "" }

type List struct {
	Kind string
	Loc  *Location
	Type Type
}

func NewList(t *List) *List { _ = "STUB: not implemented"; return nil }

func (t *List) GetKind() string { _ = "STUB: not implemented"; return "" }

func (t *List) GetLoc() *Location { _ = "STUB: not implemented"; return nil }

func (t *List) String() string { _ = "STUB: not implemented"; return "" }

type NonNull struct {
	Kind string
	Loc  *Location
	Type Type
}

func NewNonNull(t *NonNull) *NonNull { _ = "STUB: not implemented"; return nil }

func (t *NonNull) GetKind() string { _ = "STUB: not implemented"; return "" }

func (t *NonNull) GetLoc() *Location { _ = "STUB: not implemented"; return nil }

func (t *NonNull) String() string { _ = "STUB: not implemented"; return "" }
