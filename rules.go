package graphql

import (
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/visitor"
)

var SpecifiedRules = []ValidationRuleFn{
	ArgumentsOfCorrectTypeRule,
	DefaultValuesOfCorrectTypeRule,
	FieldsOnCorrectTypeRule,
	FragmentsOnCompositeTypesRule,
	KnownArgumentNamesRule,
	KnownDirectivesRule,
	KnownFragmentNamesRule,
	KnownTypeNamesRule,
	LoneAnonymousOperationRule,
	NoFragmentCyclesRule,
	NoUndefinedVariablesRule,
	NoUnusedFragmentsRule,
	NoUnusedVariablesRule,
	OverlappingFieldsCanBeMergedRule,
	PossibleFragmentSpreadsRule,
	ProvidedNonNullArgumentsRule,
	ScalarLeafsRule,
	UniqueArgumentNamesRule,
	UniqueFragmentNamesRule,
	UniqueInputFieldNamesRule,
	UniqueOperationNamesRule,
	UniqueVariableNamesRule,
	VariablesAreInputTypesRule,
	VariablesInAllowedPositionRule,
}

type ValidationRuleInstance struct {
	VisitorOpts *visitor.VisitorOptions
}
type ValidationRuleFn func(context *ValidationContext) *ValidationRuleInstance

func newValidationError(message string, nodes []ast.Node) *gqlerrors.Error {
	_ = "STUB: not implemented"
	return nil
}

func reportError(context *ValidationContext, message string, nodes []ast.Node) (string, interface{}) {
	_ = "STUB: not implemented"
	return "", nil
}

func ArgumentsOfCorrectTypeRule(context *ValidationContext) *ValidationRuleInstance {
	_ = "STUB: not implemented"
	return nil
}

func DefaultValuesOfCorrectTypeRule(context *ValidationContext) *ValidationRuleInstance {
	_ = "STUB: not implemented"
	return nil
}

func quoteStrings(slice []string) []string { _ = "STUB: not implemented"; return nil }

func quotedOrList(slice []string) string { _ = "STUB: not implemented"; return "" }

func UndefinedFieldMessage(fieldName string, ttypeName string, suggestedTypeNames []string, suggestedFieldNames []string) string {
	_ = "STUB: not implemented"
	return ""
}

func FieldsOnCorrectTypeRule(context *ValidationContext) *ValidationRuleInstance {
	_ = "STUB: not implemented"
	return nil
}

func getSuggestedTypeNames(schema *Schema, ttype Output, fieldName string) []string {
	_ = "STUB: not implemented"
	return nil
}

func getSuggestedFieldNames(schema *Schema, ttype Output, fieldName string) []string {
	_ = "STUB: not implemented"
	return nil
}

type suggestedInterface struct {
	name  string
	count int
}
type suggestedInterfaceSortedSlice []*suggestedInterface

func (s suggestedInterfaceSortedSlice) Len() int { _ = "STUB: not implemented"; return 0 }

func (s suggestedInterfaceSortedSlice) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (s suggestedInterfaceSortedSlice) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func FragmentsOnCompositeTypesRule(context *ValidationContext) *ValidationRuleInstance {
	_ = "STUB: not implemented"
	return nil
}

func unknownArgMessage(argName string, fieldName string, parentTypeName string, suggestedArgs []string) string {
	_ = "STUB: not implemented"
	return ""
}

func unknownDirectiveArgMessage(argName string, directiveName string, suggestedArgs []string) string {
	_ = "STUB: not implemented"
	return ""
}

func KnownArgumentNamesRule(context *ValidationContext) *ValidationRuleInstance {
	_ = "STUB: not implemented"
	return nil
}

func MisplaceDirectiveMessage(directiveName string, location string) string {
	_ = "STUB: not implemented"
	return ""
}

func KnownDirectivesRule(context *ValidationContext) *ValidationRuleInstance {
	_ = "STUB: not implemented"
	return nil
}

func getDirectiveLocationForASTPath(ancestors []ast.Node) string {
	_ = "STUB: not implemented"
	return ""
}

func KnownFragmentNamesRule(context *ValidationContext) *ValidationRuleInstance {
	_ = "STUB: not implemented"
	return nil
}

func unknownTypeMessage(typeName string, suggestedTypes []string) string {
	_ = "STUB: not implemented"
	return ""
}

func KnownTypeNamesRule(context *ValidationContext) *ValidationRuleInstance {
	_ = "STUB: not implemented"
	return nil
}

func LoneAnonymousOperationRule(context *ValidationContext) *ValidationRuleInstance {
	_ = "STUB: not implemented"
	return nil
}

func CycleErrorMessage(fragName string, spreadNames []string) string {
	_ = "STUB: not implemented"
	return ""
}

func NoFragmentCyclesRule(context *ValidationContext) *ValidationRuleInstance {
	_ = "STUB: not implemented"
	return nil
}

func UndefinedVarMessage(varName string, opName string) string {
	_ = "STUB: not implemented"
	return ""
}

func NoUndefinedVariablesRule(context *ValidationContext) *ValidationRuleInstance {
	_ = "STUB: not implemented"
	return nil
}

func NoUnusedFragmentsRule(context *ValidationContext) *ValidationRuleInstance {
	_ = "STUB: not implemented"
	return nil
}

func UnusedVariableMessage(varName string, opName string) string {
	_ = "STUB: not implemented"
	return ""
}

func NoUnusedVariablesRule(context *ValidationContext) *ValidationRuleInstance {
	_ = "STUB: not implemented"
	return nil
}

func getFragmentType(context *ValidationContext, name string) Type {
	_ = "STUB: not implemented"
	return *new(Type)
}

func doTypesOverlap(schema *Schema, t1 Type, t2 Type) bool { _ = "STUB: not implemented"; return false }

func PossibleFragmentSpreadsRule(context *ValidationContext) *ValidationRuleInstance {
	_ = "STUB: not implemented"
	return nil
}

func ProvidedNonNullArgumentsRule(context *ValidationContext) *ValidationRuleInstance {
	_ = "STUB: not implemented"
	return nil
}

func ScalarLeafsRule(context *ValidationContext) *ValidationRuleInstance {
	_ = "STUB: not implemented"
	return nil
}

func UniqueArgumentNamesRule(context *ValidationContext) *ValidationRuleInstance {
	_ = "STUB: not implemented"
	return nil
}

func UniqueFragmentNamesRule(context *ValidationContext) *ValidationRuleInstance {
	_ = "STUB: not implemented"
	return nil
}

func UniqueInputFieldNamesRule(context *ValidationContext) *ValidationRuleInstance {
	_ = "STUB: not implemented"
	return nil
}

func UniqueOperationNamesRule(context *ValidationContext) *ValidationRuleInstance {
	_ = "STUB: not implemented"
	return nil
}

func UniqueVariableNamesRule(context *ValidationContext) *ValidationRuleInstance {
	_ = "STUB: not implemented"
	return nil
}

func VariablesAreInputTypesRule(context *ValidationContext) *ValidationRuleInstance {
	_ = "STUB: not implemented"
	return nil
}

func effectiveType(varType Type, varDef *ast.VariableDefinition) Type {
	_ = "STUB: not implemented"
	return *new(Type)
}

func VariablesInAllowedPositionRule(context *ValidationContext) *ValidationRuleInstance {
	_ = "STUB: not implemented"
	return nil
}

func isValidLiteralValue(ttype Input, valueAST ast.Value) (bool, []string) {
	_ = "STUB: not implemented"
	return false, nil
}

type suggestionListResult struct {
	Options   []string
	Distances []float64
}

func (s suggestionListResult) Len() int { _ = "STUB: not implemented"; return 0 }

func (s suggestionListResult) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (s suggestionListResult) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func suggestionList(input string, options []string) []string { _ = "STUB: not implemented"; return nil }

func lexicalDistance(a, b string) float64 { _ = "STUB: not implemented"; return 0 }
