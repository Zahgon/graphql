package ast

type DescribableNode interface {
	GetDescription() *StringValue
}

type TypeDefinition interface {
	DescribableNode
	GetOperation() string
	GetVariableDefinitions() []*VariableDefinition
	GetSelectionSet() *SelectionSet
	GetKind() string
	GetLoc() *Location
}

var _ TypeDefinition = (*ScalarDefinition)(nil)
var _ TypeDefinition = (*ObjectDefinition)(nil)
var _ TypeDefinition = (*InterfaceDefinition)(nil)
var _ TypeDefinition = (*UnionDefinition)(nil)
var _ TypeDefinition = (*EnumDefinition)(nil)
var _ TypeDefinition = (*InputObjectDefinition)(nil)

type TypeSystemDefinition interface {
	GetOperation() string
	GetVariableDefinitions() []*VariableDefinition
	GetSelectionSet() *SelectionSet
	GetKind() string
	GetLoc() *Location
}

var _ TypeSystemDefinition = (*SchemaDefinition)(nil)
var _ TypeSystemDefinition = (TypeDefinition)(nil)
var _ TypeSystemDefinition = (*TypeExtensionDefinition)(nil)
var _ TypeSystemDefinition = (*DirectiveDefinition)(nil)

type SchemaDefinition struct {
	Kind           string
	Loc            *Location
	Directives     []*Directive
	OperationTypes []*OperationTypeDefinition
}

func NewSchemaDefinition(def *SchemaDefinition) *SchemaDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (def *SchemaDefinition) GetKind() string { _ = "STUB: not implemented"; return "" }

func (def *SchemaDefinition) GetLoc() *Location { _ = "STUB: not implemented"; return nil }

func (def *SchemaDefinition) GetVariableDefinitions() []*VariableDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (def *SchemaDefinition) GetSelectionSet() *SelectionSet { _ = "STUB: not implemented"; return nil }

func (def *SchemaDefinition) GetOperation() string { _ = "STUB: not implemented"; return "" }

type OperationTypeDefinition struct {
	Kind      string
	Loc       *Location
	Operation string
	Type      *Named
}

func NewOperationTypeDefinition(def *OperationTypeDefinition) *OperationTypeDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (def *OperationTypeDefinition) GetKind() string { _ = "STUB: not implemented"; return "" }

func (def *OperationTypeDefinition) GetLoc() *Location { _ = "STUB: not implemented"; return nil }

type ScalarDefinition struct {
	Kind        string
	Loc         *Location
	Description *StringValue
	Name        *Name
	Directives  []*Directive
}

func NewScalarDefinition(def *ScalarDefinition) *ScalarDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (def *ScalarDefinition) GetKind() string { _ = "STUB: not implemented"; return "" }

func (def *ScalarDefinition) GetLoc() *Location { _ = "STUB: not implemented"; return nil }

func (def *ScalarDefinition) GetName() *Name { _ = "STUB: not implemented"; return nil }

func (def *ScalarDefinition) GetVariableDefinitions() []*VariableDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (def *ScalarDefinition) GetSelectionSet() *SelectionSet { _ = "STUB: not implemented"; return nil }

func (def *ScalarDefinition) GetOperation() string { _ = "STUB: not implemented"; return "" }

func (def *ScalarDefinition) GetDescription() *StringValue { _ = "STUB: not implemented"; return nil }

type ObjectDefinition struct {
	Kind        string
	Loc         *Location
	Name        *Name
	Description *StringValue
	Interfaces  []*Named
	Directives  []*Directive
	Fields      []*FieldDefinition
}

func NewObjectDefinition(def *ObjectDefinition) *ObjectDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (def *ObjectDefinition) GetKind() string { _ = "STUB: not implemented"; return "" }

func (def *ObjectDefinition) GetLoc() *Location { _ = "STUB: not implemented"; return nil }

func (def *ObjectDefinition) GetName() *Name { _ = "STUB: not implemented"; return nil }

func (def *ObjectDefinition) GetVariableDefinitions() []*VariableDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (def *ObjectDefinition) GetSelectionSet() *SelectionSet { _ = "STUB: not implemented"; return nil }

func (def *ObjectDefinition) GetOperation() string { _ = "STUB: not implemented"; return "" }

func (def *ObjectDefinition) GetDescription() *StringValue { _ = "STUB: not implemented"; return nil }

type FieldDefinition struct {
	Kind        string
	Loc         *Location
	Name        *Name
	Description *StringValue
	Arguments   []*InputValueDefinition
	Type        Type
	Directives  []*Directive
}

func NewFieldDefinition(def *FieldDefinition) *FieldDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (def *FieldDefinition) GetKind() string { _ = "STUB: not implemented"; return "" }

func (def *FieldDefinition) GetLoc() *Location { _ = "STUB: not implemented"; return nil }

func (def *FieldDefinition) GetDescription() *StringValue { _ = "STUB: not implemented"; return nil }

type InputValueDefinition struct {
	Kind         string
	Loc          *Location
	Name         *Name
	Description  *StringValue
	Type         Type
	DefaultValue Value
	Directives   []*Directive
}

func NewInputValueDefinition(def *InputValueDefinition) *InputValueDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (def *InputValueDefinition) GetKind() string { _ = "STUB: not implemented"; return "" }

func (def *InputValueDefinition) GetLoc() *Location { _ = "STUB: not implemented"; return nil }

func (def *InputValueDefinition) GetDescription() *StringValue {
	_ = "STUB: not implemented"
	return nil
}

type InterfaceDefinition struct {
	Kind        string
	Loc         *Location
	Name        *Name
	Description *StringValue
	Directives  []*Directive
	Fields      []*FieldDefinition
}

func NewInterfaceDefinition(def *InterfaceDefinition) *InterfaceDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (def *InterfaceDefinition) GetKind() string { _ = "STUB: not implemented"; return "" }

func (def *InterfaceDefinition) GetLoc() *Location { _ = "STUB: not implemented"; return nil }

func (def *InterfaceDefinition) GetName() *Name { _ = "STUB: not implemented"; return nil }

func (def *InterfaceDefinition) GetVariableDefinitions() []*VariableDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (def *InterfaceDefinition) GetSelectionSet() *SelectionSet {
	_ = "STUB: not implemented"
	return nil
}

func (def *InterfaceDefinition) GetOperation() string { _ = "STUB: not implemented"; return "" }

func (def *InterfaceDefinition) GetDescription() *StringValue {
	_ = "STUB: not implemented"
	return nil
}

type UnionDefinition struct {
	Kind        string
	Loc         *Location
	Name        *Name
	Description *StringValue
	Directives  []*Directive
	Types       []*Named
}

func NewUnionDefinition(def *UnionDefinition) *UnionDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (def *UnionDefinition) GetKind() string { _ = "STUB: not implemented"; return "" }

func (def *UnionDefinition) GetLoc() *Location { _ = "STUB: not implemented"; return nil }

func (def *UnionDefinition) GetName() *Name { _ = "STUB: not implemented"; return nil }

func (def *UnionDefinition) GetVariableDefinitions() []*VariableDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (def *UnionDefinition) GetSelectionSet() *SelectionSet { _ = "STUB: not implemented"; return nil }

func (def *UnionDefinition) GetOperation() string { _ = "STUB: not implemented"; return "" }

func (def *UnionDefinition) GetDescription() *StringValue { _ = "STUB: not implemented"; return nil }

type EnumDefinition struct {
	Kind        string
	Loc         *Location
	Name        *Name
	Description *StringValue
	Directives  []*Directive
	Values      []*EnumValueDefinition
}

func NewEnumDefinition(def *EnumDefinition) *EnumDefinition { _ = "STUB: not implemented"; return nil }

func (def *EnumDefinition) GetKind() string { _ = "STUB: not implemented"; return "" }

func (def *EnumDefinition) GetLoc() *Location { _ = "STUB: not implemented"; return nil }

func (def *EnumDefinition) GetName() *Name { _ = "STUB: not implemented"; return nil }

func (def *EnumDefinition) GetVariableDefinitions() []*VariableDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (def *EnumDefinition) GetSelectionSet() *SelectionSet { _ = "STUB: not implemented"; return nil }

func (def *EnumDefinition) GetOperation() string { _ = "STUB: not implemented"; return "" }

func (def *EnumDefinition) GetDescription() *StringValue { _ = "STUB: not implemented"; return nil }

type EnumValueDefinition struct {
	Kind        string
	Loc         *Location
	Name        *Name
	Description *StringValue
	Directives  []*Directive
}

func NewEnumValueDefinition(def *EnumValueDefinition) *EnumValueDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (def *EnumValueDefinition) GetKind() string { _ = "STUB: not implemented"; return "" }

func (def *EnumValueDefinition) GetLoc() *Location { _ = "STUB: not implemented"; return nil }

func (def *EnumValueDefinition) GetDescription() *StringValue {
	_ = "STUB: not implemented"
	return nil
}

type InputObjectDefinition struct {
	Kind        string
	Loc         *Location
	Name        *Name
	Description *StringValue
	Directives  []*Directive
	Fields      []*InputValueDefinition
}

func NewInputObjectDefinition(def *InputObjectDefinition) *InputObjectDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (def *InputObjectDefinition) GetKind() string { _ = "STUB: not implemented"; return "" }

func (def *InputObjectDefinition) GetLoc() *Location { _ = "STUB: not implemented"; return nil }

func (def *InputObjectDefinition) GetName() *Name { _ = "STUB: not implemented"; return nil }

func (def *InputObjectDefinition) GetVariableDefinitions() []*VariableDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (def *InputObjectDefinition) GetSelectionSet() *SelectionSet {
	_ = "STUB: not implemented"
	return nil
}

func (def *InputObjectDefinition) GetOperation() string { _ = "STUB: not implemented"; return "" }

func (def *InputObjectDefinition) GetDescription() *StringValue {
	_ = "STUB: not implemented"
	return nil
}
