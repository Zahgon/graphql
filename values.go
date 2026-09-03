package graphql

import (
	"github.com/graphql-go/graphql/language/ast"
)

func getVariableValues(
	schema Schema,
	definitionASTs []*ast.VariableDefinition,
	inputs map[string]interface{}) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getArgumentValues(
	argDefs []*Argument, argASTs []*ast.Argument,
	variableValues map[string]interface{}) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

func getVariableValue(schema Schema, definitionAST *ast.VariableDefinition, input interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func coerceValue(ttype Input, value interface{}) interface{} { _ = "STUB: not implemented"; return nil }

func typeFromAST(schema Schema, inputTypeAST ast.Type) (Type, error) {
	_ = "STUB: not implemented"
	return *new(Type), nil
}

func isValidInputValue(value interface{}, ttype Input) (bool, []string) {
	_ = "STUB: not implemented"
	return false, nil
}

func isNullish(src interface{}) bool { _ = "STUB: not implemented"; return false }

func isIterable(src interface{}) bool { _ = "STUB: not implemented"; return false }

func valueFromAST(valueAST ast.Value, ttype Input, variables map[string]interface{}) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func invariant(condition bool, message string) error { _ = "STUB: not implemented"; return nil }

func invariantf(condition bool, format string, a ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}
