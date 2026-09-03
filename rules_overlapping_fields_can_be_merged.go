package graphql

import (
	"github.com/graphql-go/graphql/language/ast"
)

func fieldsConflictMessage(responseName string, reason conflictReason) string {
	_ = "STUB: not implemented"
	return ""
}

func fieldsConflictReasonMessage(message interface{}) string { _ = "STUB: not implemented"; return "" }

func OverlappingFieldsCanBeMergedRule(context *ValidationContext) *ValidationRuleInstance {
	_ = "STUB: not implemented"
	return nil
}

type overlappingFieldsCanBeMergedRule struct {
	context *ValidationContext

	comparedSet *pairSet

	comparedFieldsAndFragmentSet *fieldsAndFragmentSet

	cacheMap map[*ast.SelectionSet]*fieldsAndFragmentNames
}

func (rule *overlappingFieldsCanBeMergedRule) findConflictsWithinSelectionSet(parentType Named, selectionSet *ast.SelectionSet) []conflict {
	_ = "STUB: not implemented"
	return nil
}

func (rule *overlappingFieldsCanBeMergedRule) collectConflictsBetweenFieldsAndFragment(conflicts []conflict, areMutuallyExclusive bool, fieldsInfo *fieldsAndFragmentNames, fragmentName string) []conflict {
	_ = "STUB: not implemented"
	return nil
}

func (rule *overlappingFieldsCanBeMergedRule) collectConflictsBetweenFragments(conflicts []conflict, areMutuallyExclusive bool, fragmentName1 string, fragmentName2 string) []conflict {
	_ = "STUB: not implemented"
	return nil
}

func (rule *overlappingFieldsCanBeMergedRule) findConflictsBetweenSubSelectionSets(areMutuallyExclusive bool, parentType1 Named, selectionSet1 *ast.SelectionSet, parentType2 Named, selectionSet2 *ast.SelectionSet) []conflict {
	_ = "STUB: not implemented"
	return nil
}

func (rule *overlappingFieldsCanBeMergedRule) collectConflictsWithin(conflicts []conflict, fieldsInfo *fieldsAndFragmentNames) []conflict {
	_ = "STUB: not implemented"
	return nil
}

func (rule *overlappingFieldsCanBeMergedRule) collectConflictsBetween(conflicts []conflict, parentFieldsAreMutuallyExclusive bool,
	fieldsInfo1 *fieldsAndFragmentNames,
	fieldsInfo2 *fieldsAndFragmentNames) []conflict {
	_ = "STUB: not implemented"
	return nil
}

func (rule *overlappingFieldsCanBeMergedRule) findConflict(parentFieldsAreMutuallyExclusive bool, responseName string, field *fieldDefPair, field2 *fieldDefPair) *conflict {
	_ = "STUB: not implemented"
	return nil
}

func (rule *overlappingFieldsCanBeMergedRule) getFieldsAndFragmentNames(parentType Named, selectionSet *ast.SelectionSet) *fieldsAndFragmentNames {
	_ = "STUB: not implemented"
	return nil
}

func (rule *overlappingFieldsCanBeMergedRule) getReferencedFieldsAndFragmentNames(fragment *ast.FragmentDefinition) *fieldsAndFragmentNames {
	_ = "STUB: not implemented"
	return nil
}

type conflictReason struct {
	Name    string
	Message interface{}
}
type conflict struct {
	Reason      conflictReason
	FieldsLeft  []ast.Node
	FieldsRight []ast.Node
}

type fieldDefPair struct {
	ParentType Named
	Field      *ast.Field
	FieldDef   *FieldDefinition
}
type astAndDefCollection map[string][]*fieldDefPair

type fieldsAndFragmentNames struct {
	fieldMap      astAndDefCollection
	fieldsOrder   []string
	fragmentNames []string
}

type fieldsAndFragmentSet struct {
	data map[*fieldsAndFragmentNames]map[string]bool
}

func newFieldsAndFragmentSet() *fieldsAndFragmentSet { _ = "STUB: not implemented"; return nil }

func (s *fieldsAndFragmentSet) Has(fields *fieldsAndFragmentNames, fragmentName string, areMutuallyExclusive bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *fieldsAndFragmentSet) Add(fields *fieldsAndFragmentNames, fragmentName string, areMutuallyExclusive bool) {
	_ = "STUB: not implemented"
	return
}

type pairSet struct {
	data map[string]map[string]bool
}

func newPairSet() *pairSet { _ = "STUB: not implemented"; return nil }

func (pair *pairSet) Has(a string, b string, areMutuallyExclusive bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (pair *pairSet) Add(a string, b string, areMutuallyExclusive bool) {
	_ = "STUB: not implemented"
	return
}

func pairSetAdd(data map[string]map[string]bool, a, b string, areMutuallyExclusive bool) map[string]map[string]bool {
	_ = "STUB: not implemented"
	return nil
}

func sameArguments(args1 []*ast.Argument, args2 []*ast.Argument) bool {
	_ = "STUB: not implemented"
	return false
}

func sameValue(value1 ast.Value, value2 ast.Value) bool { _ = "STUB: not implemented"; return false }

func doTypesConflict(type1 Output, type2 Output) bool { _ = "STUB: not implemented"; return false }

func subfieldConflicts(conflicts []conflict, responseName string, ast1 *ast.Field, ast2 *ast.Field) *conflict {
	_ = "STUB: not implemented"
	return nil
}
