package graphql

import (
	"github.com/graphql-go/graphql/language/ast"
)

type fieldDefFn func(schema *Schema, parentType Type, fieldAST *ast.Field) *FieldDefinition

type TypeInfo struct {
	schema          *Schema
	typeStack       []Output
	parentTypeStack []Composite
	inputTypeStack  []Input
	fieldDefStack   []*FieldDefinition
	directive       *Directive
	argument        *Argument
	getFieldDef     fieldDefFn
}

type TypeInfoConfig struct {
	Schema *Schema

	FieldDefFn fieldDefFn
}

func NewTypeInfo(opts *TypeInfoConfig) *TypeInfo { _ = "STUB: not implemented"; return nil }

func (ti *TypeInfo) Type() Output { _ = "STUB: not implemented"; return *new(Output) }

func (ti *TypeInfo) ParentType() Composite { _ = "STUB: not implemented"; return *new(Composite) }

func (ti *TypeInfo) InputType() Input { _ = "STUB: not implemented"; return *new(Input) }

func (ti *TypeInfo) FieldDef() *FieldDefinition { _ = "STUB: not implemented"; return nil }

func (ti *TypeInfo) Directive() *Directive { _ = "STUB: not implemented"; return nil }

func (ti *TypeInfo) Argument() *Argument { _ = "STUB: not implemented"; return nil }

func (ti *TypeInfo) Enter(node ast.Node) { _ = "STUB: not implemented"; return }

func (ti *TypeInfo) Leave(node ast.Node) { _ = "STUB: not implemented"; return }

func DefaultTypeInfoFieldDef(schema *Schema, parentType Type, fieldAST *ast.Field) *FieldDefinition {
	_ = "STUB: not implemented"
	return nil
}
