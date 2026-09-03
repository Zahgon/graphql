package graphql

import (
	"github.com/graphql-go/graphql/language/ast"
)

func normalizeDocument(schema *Schema, doc *ast.Document, operationName string) (*ast.Document, map[string]interface{}, string, error) {
	_ = "STUB: not implemented"
	return nil, nil, "", nil
}

func fingerprintDocument(doc *ast.Document, op *ast.OperationDefinition, operationName string) string {
	_ = "STUB: not implemented"
	return ""
}

func collectFragmentDefs(doc *ast.Document) map[string]*ast.FragmentDefinition {
	_ = "STUB: not implemented"
	return nil
}

type fingerprintWriter struct {
	h         interface{ Write([]byte) (int, error) }
	fragments map[string]*ast.FragmentDefinition
	visited   map[string]bool
}

func (w *fingerprintWriter) writeString(s string) { _ = "STUB: not implemented"; return }
func (w *fingerprintWriter) writeByte(b byte)     { _ = "STUB: not implemented"; return }

func (w *fingerprintWriter) writeVariableDefs(defs []*ast.VariableDefinition) {
	_ = "STUB: not implemented"
	return
}

func (w *fingerprintWriter) writeType(t ast.Type) { _ = "STUB: not implemented"; return }

func (w *fingerprintWriter) writeSelectionSet(sel *ast.SelectionSet) {
	_ = "STUB: not implemented"
	return
}

func (w *fingerprintWriter) writeFragmentBody(name string) { _ = "STUB: not implemented"; return }

func (w *fingerprintWriter) writeValue(v ast.Value) { _ = "STUB: not implemented"; return }

type normCtx struct {
	schema     *Schema
	counter    int
	synthArgs  map[string]interface{}
	newVarDefs []*ast.VariableDefinition
}

func (c *normCtx) nextName() string { _ = "STUB: not implemented"; return "" }

func (c *normCtx) normalizeSelectionSet(sel *ast.SelectionSet, parentType *Object) {
	_ = "STUB: not implemented"
	return
}

func (c *normCtx) normalizeField(f *ast.Field, parentType *Object) {
	_ = "STUB: not implemented"
	return
}

func (c *normCtx) tryExtract(value ast.Value, expected Input) (ast.Value, bool) {
	_ = "STUB: not implemented"
	return *new(ast.Value), false
}

func typeASTFromGoType(t Input) ast.Type { _ = "STUB: not implemented"; return *new(ast.Type) }

func unwrapToNamed(t Type) Type { _ = "STUB: not implemented"; return *new(Type) }

func cloneOperation(op *ast.OperationDefinition) *ast.OperationDefinition {
	_ = "STUB: not implemented"
	return nil
}

func cloneSelectionSet(sel *ast.SelectionSet) *ast.SelectionSet {
	_ = "STUB: not implemented"
	return nil
}

func cloneField(f *ast.Field) *ast.Field { _ = "STUB: not implemented"; return nil }
