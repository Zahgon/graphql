package graphql

import (
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/graphql-go/graphql/language/ast"
)

type ValidationResult struct {
	IsValid bool
	Errors  []gqlerrors.FormattedError
}

func ValidateDocument(schema *Schema, astDoc *ast.Document, rules []ValidationRuleFn) (vr ValidationResult) {
	_ = "STUB: not implemented"
	return *new(ValidationResult)
}

func VisitUsingRules(schema *Schema, typeInfo *TypeInfo, astDoc *ast.Document, rules []ValidationRuleFn) []gqlerrors.FormattedError {
	_ = "STUB: not implemented"
	return nil
}

type HasSelectionSet interface {
	GetKind() string
	GetLoc() *ast.Location
	GetSelectionSet() *ast.SelectionSet
}

var _ HasSelectionSet = (*ast.OperationDefinition)(nil)
var _ HasSelectionSet = (*ast.FragmentDefinition)(nil)

type VariableUsage struct {
	Node *ast.Variable
	Type Input
}

type ValidationContext struct {
	schema                         *Schema
	astDoc                         *ast.Document
	typeInfo                       *TypeInfo
	errors                         []gqlerrors.FormattedError
	fragments                      map[string]*ast.FragmentDefinition
	variableUsages                 map[HasSelectionSet][]*VariableUsage
	recursiveVariableUsages        map[*ast.OperationDefinition][]*VariableUsage
	recursivelyReferencedFragments map[*ast.OperationDefinition][]*ast.FragmentDefinition
	fragmentSpreads                map[*ast.SelectionSet][]*ast.FragmentSpread
}

func NewValidationContext(schema *Schema, astDoc *ast.Document, typeInfo *TypeInfo) *ValidationContext {
	_ = "STUB: not implemented"
	return nil
}

func (ctx *ValidationContext) ReportError(err error) { _ = "STUB: not implemented"; return }

func (ctx *ValidationContext) Errors() []gqlerrors.FormattedError {
	_ = "STUB: not implemented"
	return nil
}

func (ctx *ValidationContext) Schema() *Schema { _ = "STUB: not implemented"; return nil }

func (ctx *ValidationContext) Document() *ast.Document { _ = "STUB: not implemented"; return nil }

func (ctx *ValidationContext) Fragment(name string) *ast.FragmentDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (ctx *ValidationContext) FragmentSpreads(node *ast.SelectionSet) []*ast.FragmentSpread {
	_ = "STUB: not implemented"
	return nil
}

func (ctx *ValidationContext) RecursivelyReferencedFragments(operation *ast.OperationDefinition) []*ast.FragmentDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (ctx *ValidationContext) VariableUsages(node HasSelectionSet) []*VariableUsage {
	_ = "STUB: not implemented"
	return nil
}

func (ctx *ValidationContext) RecursiveVariableUsages(operation *ast.OperationDefinition) []*VariableUsage {
	_ = "STUB: not implemented"
	return nil
}

func (ctx *ValidationContext) Type() Output { _ = "STUB: not implemented"; return *new(Output) }

func (ctx *ValidationContext) ParentType() Composite {
	_ = "STUB: not implemented"
	return *new(Composite)
}

func (ctx *ValidationContext) InputType() Input { _ = "STUB: not implemented"; return *new(Input) }

func (ctx *ValidationContext) FieldDef() *FieldDefinition { _ = "STUB: not implemented"; return nil }

func (ctx *ValidationContext) Directive() *Directive { _ = "STUB: not implemented"; return nil }

func (ctx *ValidationContext) Argument() *Argument { _ = "STUB: not implemented"; return nil }
