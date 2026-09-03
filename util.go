package graphql

import (
	"reflect"
)

const TAG = "json"

func BindFields(obj interface{}) Fields { _ = "STUB: not implemented"; return *new(Fields) }

func getGraphType(tipe reflect.Type) Output { _ = "STUB: not implemented"; return *new(Output) }

func getGraphList(tipe reflect.Type) *List { _ = "STUB: not implemented"; return nil }

func appendFields(dest, origin Fields) Fields { _ = "STUB: not implemented"; return *new(Fields) }

func extractValue(originTag string, obj interface{}) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func extractTag(tag reflect.StructTag) string { _ = "STUB: not implemented"; return "" }

func BindArg(obj interface{}, tags ...string) FieldConfigArgument {
	_ = "STUB: not implemented"
	return *new(FieldConfigArgument)
}

func inArray(slice interface{}, item interface{}) bool { _ = "STUB: not implemented"; return false }
