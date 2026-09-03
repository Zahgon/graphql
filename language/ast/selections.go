package ast

type Selection interface {
	GetSelectionSet() *SelectionSet
}

var _ Selection = (*Field)(nil)
var _ Selection = (*FragmentSpread)(nil)
var _ Selection = (*InlineFragment)(nil)

type Field struct {
	Kind         string
	Loc          *Location
	Alias        *Name
	Name         *Name
	Arguments    []*Argument
	Directives   []*Directive
	SelectionSet *SelectionSet
}

func NewField(f *Field) *Field { _ = "STUB: not implemented"; return nil }

func (f *Field) GetKind() string { _ = "STUB: not implemented"; return "" }

func (f *Field) GetLoc() *Location { _ = "STUB: not implemented"; return nil }

func (f *Field) GetSelectionSet() *SelectionSet { _ = "STUB: not implemented"; return nil }

type FragmentSpread struct {
	Kind       string
	Loc        *Location
	Name       *Name
	Directives []*Directive
}

func NewFragmentSpread(fs *FragmentSpread) *FragmentSpread { _ = "STUB: not implemented"; return nil }

func (fs *FragmentSpread) GetKind() string { _ = "STUB: not implemented"; return "" }

func (fs *FragmentSpread) GetLoc() *Location { _ = "STUB: not implemented"; return nil }

func (fs *FragmentSpread) GetSelectionSet() *SelectionSet { _ = "STUB: not implemented"; return nil }

type InlineFragment struct {
	Kind          string
	Loc           *Location
	TypeCondition *Named
	Directives    []*Directive
	SelectionSet  *SelectionSet
}

func NewInlineFragment(f *InlineFragment) *InlineFragment { _ = "STUB: not implemented"; return nil }

func (f *InlineFragment) GetKind() string { _ = "STUB: not implemented"; return "" }

func (f *InlineFragment) GetLoc() *Location { _ = "STUB: not implemented"; return nil }

func (f *InlineFragment) GetSelectionSet() *SelectionSet { _ = "STUB: not implemented"; return nil }

type SelectionSet struct {
	Kind       string
	Loc        *Location
	Selections []Selection
}

func NewSelectionSet(ss *SelectionSet) *SelectionSet { _ = "STUB: not implemented"; return nil }

func (ss *SelectionSet) GetKind() string { _ = "STUB: not implemented"; return "" }

func (ss *SelectionSet) GetLoc() *Location { _ = "STUB: not implemented"; return nil }
