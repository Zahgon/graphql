package ast

type Definition interface {
	GetOperation() string
	GetVariableDefinitions() []*VariableDefinition
	GetSelectionSet() *SelectionSet
	GetKind() string
	GetLoc() *Location
}

var _ Definition = (*OperationDefinition)(nil)
var _ Definition = (*FragmentDefinition)(nil)
var _ Definition = (TypeSystemDefinition)(nil)

const (
	OperationTypeQuery        = "query"
	OperationTypeMutation     = "mutation"
	OperationTypeSubscription = "subscription"
)

type OperationDefinition struct {
	Kind                string
	Loc                 *Location
	Operation           string
	Name                *Name
	VariableDefinitions []*VariableDefinition
	Directives          []*Directive
	SelectionSet        *SelectionSet
}

func NewOperationDefinition(op *OperationDefinition) *OperationDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (op *OperationDefinition) GetKind() string { _ = "STUB: not implemented"; return "" }

func (op *OperationDefinition) GetLoc() *Location { _ = "STUB: not implemented"; return nil }

func (op *OperationDefinition) GetOperation() string { _ = "STUB: not implemented"; return "" }

func (op *OperationDefinition) GetName() *Name { _ = "STUB: not implemented"; return nil }

func (op *OperationDefinition) GetVariableDefinitions() []*VariableDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (op *OperationDefinition) GetDirectives() []*Directive { _ = "STUB: not implemented"; return nil }

func (op *OperationDefinition) GetSelectionSet() *SelectionSet {
	_ = "STUB: not implemented"
	return nil
}

type FragmentDefinition struct {
	Kind                string
	Loc                 *Location
	Operation           string
	Name                *Name
	VariableDefinitions []*VariableDefinition
	TypeCondition       *Named
	Directives          []*Directive
	SelectionSet        *SelectionSet
}

func NewFragmentDefinition(fd *FragmentDefinition) *FragmentDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (fd *FragmentDefinition) GetKind() string { _ = "STUB: not implemented"; return "" }

func (fd *FragmentDefinition) GetLoc() *Location { _ = "STUB: not implemented"; return nil }

func (fd *FragmentDefinition) GetOperation() string { _ = "STUB: not implemented"; return "" }

func (fd *FragmentDefinition) GetName() *Name { _ = "STUB: not implemented"; return nil }

func (fd *FragmentDefinition) GetVariableDefinitions() []*VariableDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (fd *FragmentDefinition) GetSelectionSet() *SelectionSet {
	_ = "STUB: not implemented"
	return nil
}

type VariableDefinition struct {
	Kind         string
	Loc          *Location
	Variable     *Variable
	Type         Type
	DefaultValue Value
}

func NewVariableDefinition(vd *VariableDefinition) *VariableDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (vd *VariableDefinition) GetKind() string { _ = "STUB: not implemented"; return "" }

func (vd *VariableDefinition) GetLoc() *Location { _ = "STUB: not implemented"; return nil }

type TypeExtensionDefinition struct {
	Kind       string
	Loc        *Location
	Definition *ObjectDefinition
}

func NewTypeExtensionDefinition(def *TypeExtensionDefinition) *TypeExtensionDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (def *TypeExtensionDefinition) GetKind() string { _ = "STUB: not implemented"; return "" }

func (def *TypeExtensionDefinition) GetLoc() *Location { _ = "STUB: not implemented"; return nil }

func (def *TypeExtensionDefinition) GetVariableDefinitions() []*VariableDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (def *TypeExtensionDefinition) GetSelectionSet() *SelectionSet {
	_ = "STUB: not implemented"
	return nil
}

func (def *TypeExtensionDefinition) GetOperation() string { _ = "STUB: not implemented"; return "" }

type DirectiveDefinition struct {
	Kind        string
	Loc         *Location
	Name        *Name
	Description *StringValue
	Arguments   []*InputValueDefinition
	Locations   []*Name
}

func NewDirectiveDefinition(def *DirectiveDefinition) *DirectiveDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (def *DirectiveDefinition) GetKind() string { _ = "STUB: not implemented"; return "" }

func (def *DirectiveDefinition) GetLoc() *Location { _ = "STUB: not implemented"; return nil }

func (def *DirectiveDefinition) GetVariableDefinitions() []*VariableDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (def *DirectiveDefinition) GetSelectionSet() *SelectionSet {
	_ = "STUB: not implemented"
	return nil
}

func (def *DirectiveDefinition) GetOperation() string { _ = "STUB: not implemented"; return "" }

func (def *DirectiveDefinition) GetDescription() *StringValue {
	_ = "STUB: not implemented"
	return nil
}
