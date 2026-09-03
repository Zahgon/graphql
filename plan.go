package graphql

import (
	"sync"

	"github.com/graphql-go/graphql/language/ast"
)

type Plan struct {
	schema     *Schema
	operation  *ast.OperationDefinition
	fragments  map[string]ast.Definition
	rootType   *Object
	root       *selectionPlan
	isMutation bool

	abstractMu sync.Mutex
}

type selectionPlan struct {
	parentType *Object
	fields     []*fieldPlan
}

type fieldPlan struct {
	responseKey string
	fieldName   string
	fieldDef    *FieldDefinition
	fieldASTs   []*ast.Field
	args        argPlan
	returnType  Output

	skipPredicate func(map[string]interface{}) bool

	sub                  *selectionPlan
	abstractAlternatives map[*Object]*selectionPlan
}

type argPlan struct {
	static map[string]interface{}

	hasVariables bool

	fieldDefArgs []*Argument
	argASTs      []*ast.Argument
}

func PlanQuery(schema *Schema, doc *ast.Document, operationName string) (*Plan, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Plan) planSelectionSet(parentType *Object, selectionSet *ast.SelectionSet, visitedFragmentNames map[string]bool) *selectionPlan {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plan) planMergedFieldChildren(fp *fieldPlan) { _ = "STUB: not implemented"; return }

func (p *Plan) abstractAlternative(fp *fieldPlan, runtimeType *Object) *selectionPlan {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plan) planMergedSelectionsForType(parentType *Object, fieldASTs []*ast.Field) *selectionPlan {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plan) collectInto(parentType *Object, selectionSet *ast.SelectionSet, visitedFragmentNames map[string]bool, sp *selectionPlan, keyed map[string]int, parentPred func(map[string]interface{}) bool) {
	_ = "STUB: not implemented"
	return
}

func andPredicates(a, b func(map[string]interface{}) bool) func(map[string]interface{}) bool {
	_ = "STUB: not implemented"
	return nil
}

func unwrapNamedType(t Output) Output { _ = "STUB: not implemented"; return *new(Output) }

func planArguments(argDefs []*Argument, argASTs []*ast.Argument) argPlan {
	_ = "STUB: not implemented"
	return *new(argPlan)
}

func astHasVariables(argASTs []*ast.Argument) bool { _ = "STUB: not implemented"; return false }

func valueHasVariables(v ast.Value) bool { _ = "STUB: not implemented"; return false }

func planDirectives(directives []*ast.Directive) (pred func(map[string]interface{}) bool, alwaysSkip bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func planFragmentMatches(schema Schema, typeConditionAST *ast.Named, runtime *Object) bool {
	_ = "STUB: not implemented"
	return false
}

func ExecutePlan(plan *Plan, p ExecuteParams) (result *Result) {
	_ = "STUB: not implemented"
	return nil
}

func executePlannedSelection(eCtx *executionContext, sp *selectionPlan, source interface{}, parentType *Object, path *ResponsePath) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

func resolvePlannedField(eCtx *executionContext, parentType *Object, source interface{}, fp *fieldPlan, path *ResponsePath) (result interface{}, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func completePlannedValueCatchingError(eCtx *executionContext, returnType Type, fp *fieldPlan, info ResolveInfo, path *ResponsePath, result interface{}) (completed interface{}) {
	_ = "STUB: not implemented"
	return nil
}

func completePlannedValue(eCtx *executionContext, returnType Type, fp *fieldPlan, info ResolveInfo, path *ResponsePath, result interface{}) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func completePlannedThunkValueCatchingError(eCtx *executionContext, returnType Type, fp *fieldPlan, info ResolveInfo, path *ResponsePath, result interface{}) (completed interface{}) {
	_ = "STUB: not implemented"
	return nil
}

func completePlannedListValue(eCtx *executionContext, returnType *List, fp *fieldPlan, info ResolveInfo, path *ResponsePath, result interface{}) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func completePlannedObjectValue(eCtx *executionContext, returnType *Object, fp *fieldPlan, info ResolveInfo, path *ResponsePath, result interface{}) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func completePlannedAbstractValue(eCtx *executionContext, returnType Abstract, fp *fieldPlan, info ResolveInfo, path *ResponsePath, result interface{}) interface{} {
	_ = "STUB: not implemented"
	return nil
}
