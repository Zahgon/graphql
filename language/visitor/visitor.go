package visitor

import (
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/typeInfo"
)

const (
	ActionNoChange = ""
	ActionBreak    = "BREAK"
	ActionSkip     = "SKIP"
	ActionUpdate   = "UPDATE"
)

type KeyMap map[string][]string

var QueryDocumentKeys = KeyMap{
	"Name":     []string{},
	"Document": []string{"Definitions"},
	"OperationDefinition": []string{
		"Name",
		"VariableDefinitions",
		"Directives",
		"SelectionSet",
	},
	"VariableDefinition": []string{
		"Variable",
		"Type",
		"DefaultValue",
	},
	"Variable":     []string{"Name"},
	"SelectionSet": []string{"Selections"},
	"Field": []string{
		"Alias",
		"Name",
		"Arguments",
		"Directives",
		"SelectionSet",
	},
	"Argument": []string{
		"Name",
		"Value",
	},

	"FragmentSpread": []string{
		"Name",
		"Directives",
	},
	"InlineFragment": []string{
		"TypeCondition",
		"Directives",
		"SelectionSet",
	},
	"FragmentDefinition": []string{
		"Name",
		"TypeCondition",
		"Directives",
		"SelectionSet",
	},

	"IntValue":     []string{},
	"FloatValue":   []string{},
	"StringValue":  []string{},
	"BooleanValue": []string{},
	"EnumValue":    []string{},
	"ListValue":    []string{"Values"},
	"ObjectValue":  []string{"Fields"},
	"ObjectField": []string{
		"Name",
		"Value",
	},

	"Directive": []string{
		"Name",
		"Arguments",
	},

	"Named":   []string{"Name"},
	"List":    []string{"Type"},
	"NonNull": []string{"Type"},

	"SchemaDefinition": []string{
		"Directives",
		"OperationTypes",
	},
	"OperationTypeDefinition": []string{"Type"},

	"ScalarDefinition": []string{
		"Name",
		"Directives",
	},
	"ObjectDefinition": []string{
		"Name",
		"Interfaces",
		"Directives",
		"Fields",
	},
	"FieldDefinition": []string{
		"Name",
		"Arguments",
		"Type",
		"Directives",
	},
	"InputValueDefinition": []string{
		"Name",
		"Type",
		"DefaultValue",
		"Directives",
	},
	"InterfaceDefinition": []string{
		"Name",
		"Directives",
		"Fields",
	},
	"UnionDefinition": []string{
		"Name",
		"Directives",
		"Types",
	},
	"EnumDefinition": []string{
		"Name",
		"Directives",
		"Values",
	},
	"EnumValueDefinition": []string{
		"Name",
		"Directives",
	},
	"InputObjectDefinition": []string{
		"Name",
		"Directives",
		"Fields",
	},

	"TypeExtensionDefinition": []string{"Definition"},

	"DirectiveDefinition": []string{"Name", "Arguments", "Locations"},
}

type stack struct {
	Index   int
	Keys    []interface{}
	Edits   []*edit
	inSlice bool
	Prev    *stack
}
type edit struct {
	Key   interface{}
	Value interface{}
}

type VisitFuncParams struct {
	Node      interface{}
	Key       interface{}
	Parent    ast.Node
	Path      []interface{}
	Ancestors []ast.Node
}

type VisitFunc func(p VisitFuncParams) (string, interface{})

type NamedVisitFuncs struct {
	Kind  VisitFunc
	Leave VisitFunc
	Enter VisitFunc
}

type VisitorOptions struct {
	KindFuncMap map[string]NamedVisitFuncs
	Enter       VisitFunc
	Leave       VisitFunc

	EnterKindMap map[string]VisitFunc
	LeaveKindMap map[string]VisitFunc
}

func Visit(root ast.Node, visitorOpts *VisitorOptions, keyMap KeyMap) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func pop(a []interface{}) (interface{}, []interface{}) { _ = "STUB: not implemented"; return nil, nil }

func popNodeSlice(a [][]interface{}) ([]interface{}, [][]interface{}) {
	_ = "STUB: not implemented"
	return nil, nil
}

func removeNodeByIndex(a []interface{}, pos int) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

func convertMap(src interface{}) (dest map[string]interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getFieldValue(obj interface{}, key interface{}) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func updateNodeField(src interface{}, targetName string, target interface{}) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func toSliceInterfaces(src interface{}) []interface{} { _ = "STUB: not implemented"; return nil }

func isSlice(value interface{}) bool { _ = "STUB: not implemented"; return false }

func isStructNode(node interface{}) bool { _ = "STUB: not implemented"; return false }

func isNode(node interface{}) bool { _ = "STUB: not implemented"; return false }

func isNilNode(node interface{}) bool { _ = "STUB: not implemented"; return false }

func VisitInParallel(visitorOptsSlice ...*VisitorOptions) *VisitorOptions {
	_ = "STUB: not implemented"
	return nil
}

func VisitWithTypeInfo(ttypeInfo typeInfo.TypeInfoI, visitorOpts *VisitorOptions) *VisitorOptions {
	_ = "STUB: not implemented"
	return nil
}

func GetVisitFn(visitorOpts *VisitorOptions, kind string, isLeaving bool) VisitFunc {
	_ = "STUB: not implemented"
	return *new(VisitFunc)
}
