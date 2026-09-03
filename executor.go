package graphql

import (
	"context"

	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/graphql-go/graphql/language/ast"
)

type ExecuteParams struct {
	Schema        Schema
	Root          interface{}
	AST           *ast.Document
	OperationName string
	Args          map[string]interface{}

	Context context.Context
}

func Execute(p ExecuteParams) (result *Result) { _ = "STUB: not implemented"; return nil }

type buildExecutionCtxParams struct {
	Schema        Schema
	Root          interface{}
	AST           *ast.Document
	OperationName string
	Args          map[string]interface{}
	Result        *Result
	Context       context.Context
}

type executionContext struct {
	Schema         Schema
	Fragments      map[string]ast.Definition
	Root           interface{}
	Operation      ast.Definition
	VariableValues map[string]interface{}
	Errors         []gqlerrors.FormattedError
	Context        context.Context

	plan *Plan
}

func buildExecutionContext(p buildExecutionCtxParams) (*executionContext, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getOperationRootType(schema Schema, operation ast.Definition) (*Object, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type dethunkQueue struct {
	DethunkFuncs []func()
}

func (d *dethunkQueue) push(f func()) { _ = "STUB: not implemented"; return }

func (d *dethunkQueue) shift() func() { _ = "STUB: not implemented"; return nil }

func dethunkMapWithBreadthFirstTraversal(finalResults map[string]interface{}) {
	_ = "STUB: not implemented"
	return
}

func dethunkMapBreadthFirst(m map[string]interface{}, dethunkQueue *dethunkQueue) {
	_ = "STUB: not implemented"
	return
}

func dethunkListBreadthFirst(list []interface{}, dethunkQueue *dethunkQueue) {
	_ = "STUB: not implemented"
	return
}

func dethunkMapDepthFirst(m map[string]interface{}) { _ = "STUB: not implemented"; return }

func dethunkListDepthFirst(list []interface{}) { _ = "STUB: not implemented"; return }

type collectFieldsParams struct {
	ExeContext           *executionContext
	RuntimeType          *Object
	SelectionSet         *ast.SelectionSet
	Fields               map[string][]*ast.Field
	VisitedFragmentNames map[string]bool
}

func collectFields(p collectFieldsParams) (fields map[string][]*ast.Field) {
	_ = "STUB: not implemented"
	return nil
}

func shouldIncludeNode(eCtx *executionContext, directives []*ast.Directive) bool {
	_ = "STUB: not implemented"
	return false
}

func doesFragmentConditionMatch(eCtx *executionContext, fragment ast.Node, ttype *Object) bool {
	_ = "STUB: not implemented"
	return false
}

func getFieldEntryKey(node *ast.Field) string { _ = "STUB: not implemented"; return "" }

func handleFieldError(r interface{}, fieldNodes []ast.Node, path *ResponsePath, returnType Output, eCtx *executionContext) {
	_ = "STUB: not implemented"
	return
}

func completeLeafValue(returnType Leaf, result interface{}) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func defaultResolveTypeFn(p ResolveTypeParams, abstractType Abstract) *Object {
	_ = "STUB: not implemented"
	return nil
}

type FieldResolver interface {
	Resolve(p ResolveParams) (interface{}, error)
}

func DefaultResolveFn(p ResolveParams) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getFieldDef(schema Schema, parentType *Object, fieldName string) *FieldDefinition {
	_ = "STUB: not implemented"
	return nil
}
