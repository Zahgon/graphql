package graphql

import (
	"context"
	"regexp"

	"github.com/graphql-go/graphql/language/ast"
)

type Type interface {
	Name() string
	Description() string
	String() string
	Error() error
}

var _ Type = (*Scalar)(nil)
var _ Type = (*Object)(nil)
var _ Type = (*Interface)(nil)
var _ Type = (*Union)(nil)
var _ Type = (*Enum)(nil)
var _ Type = (*InputObject)(nil)
var _ Type = (*List)(nil)
var _ Type = (*NonNull)(nil)
var _ Type = (*Argument)(nil)

type Input interface {
	Name() string
	Description() string
	String() string
	Error() error
}

var _ Input = (*Scalar)(nil)
var _ Input = (*Enum)(nil)
var _ Input = (*InputObject)(nil)
var _ Input = (*List)(nil)
var _ Input = (*NonNull)(nil)

func IsInputType(ttype Type) bool { _ = "STUB: not implemented"; return false }

func IsOutputType(ttype Type) bool { _ = "STUB: not implemented"; return false }

type Leaf interface {
	Name() string
	Description() string
	String() string
	Error() error
	Serialize(value interface{}) interface{}
}

var _ Leaf = (*Scalar)(nil)
var _ Leaf = (*Enum)(nil)

func IsLeafType(ttype Type) bool { _ = "STUB: not implemented"; return false }

type Output interface {
	Name() string
	Description() string
	String() string
	Error() error
}

var _ Output = (*Scalar)(nil)
var _ Output = (*Object)(nil)
var _ Output = (*Interface)(nil)
var _ Output = (*Union)(nil)
var _ Output = (*Enum)(nil)
var _ Output = (*List)(nil)
var _ Output = (*NonNull)(nil)

type Composite interface {
	Name() string
	Description() string
	String() string
	Error() error
}

var _ Composite = (*Object)(nil)
var _ Composite = (*Interface)(nil)
var _ Composite = (*Union)(nil)

func IsCompositeType(ttype interface{}) bool { _ = "STUB: not implemented"; return false }

type Abstract interface {
	Name() string
}

var _ Abstract = (*Interface)(nil)
var _ Abstract = (*Union)(nil)

func IsAbstractType(ttype interface{}) bool { _ = "STUB: not implemented"; return false }

type Nullable interface {
}

var _ Nullable = (*Scalar)(nil)
var _ Nullable = (*Object)(nil)
var _ Nullable = (*Interface)(nil)
var _ Nullable = (*Union)(nil)
var _ Nullable = (*Enum)(nil)
var _ Nullable = (*InputObject)(nil)
var _ Nullable = (*List)(nil)

func GetNullable(ttype Type) Nullable { _ = "STUB: not implemented"; return *new(Nullable) }

type Named interface {
	String() string
}

var _ Named = (*Scalar)(nil)
var _ Named = (*Object)(nil)
var _ Named = (*Interface)(nil)
var _ Named = (*Union)(nil)
var _ Named = (*Enum)(nil)
var _ Named = (*InputObject)(nil)

func GetNamed(ttype Type) Named { _ = "STUB: not implemented"; return *new(Named) }

type Scalar struct {
	PrivateName        string `json:"name"`
	PrivateDescription string `json:"description"`

	scalarConfig ScalarConfig
	err          error
}

type SerializeFn func(value interface{}) interface{}

type ParseValueFn func(value interface{}) interface{}

type ParseLiteralFn func(valueAST ast.Value) interface{}

type ScalarConfig struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Serialize    SerializeFn
	ParseValue   ParseValueFn
	ParseLiteral ParseLiteralFn
}

func NewScalar(config ScalarConfig) *Scalar { _ = "STUB: not implemented"; return nil }

func (st *Scalar) Serialize(value interface{}) interface{} { _ = "STUB: not implemented"; return nil }

func (st *Scalar) ParseValue(value interface{}) interface{} { _ = "STUB: not implemented"; return nil }

func (st *Scalar) ParseLiteral(valueAST ast.Value) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (st *Scalar) Name() string { _ = "STUB: not implemented"; return "" }

func (st *Scalar) Description() string { _ = "STUB: not implemented"; return "" }

func (st *Scalar) String() string { _ = "STUB: not implemented"; return "" }

func (st *Scalar) Error() error { _ = "STUB: not implemented"; return nil }

type Object struct {
	PrivateName        string `json:"name"`
	PrivateDescription string `json:"description"`
	IsTypeOf           IsTypeOfFn

	typeConfig            ObjectConfig
	initialisedFields     bool
	fields                FieldDefinitionMap
	initialisedInterfaces bool
	interfaces            []*Interface

	err error
}

type IsTypeOfParams struct {
	Value interface{}

	Info ResolveInfo

	Context context.Context
}

type IsTypeOfFn func(p IsTypeOfParams) bool

type InterfacesThunk func() []*Interface

type ObjectConfig struct {
	Name        string      `json:"name"`
	Interfaces  interface{} `json:"interfaces"`
	Fields      interface{} `json:"fields"`
	IsTypeOf    IsTypeOfFn  `json:"isTypeOf"`
	Description string      `json:"description"`
}

type FieldsThunk func() Fields

func NewObject(config ObjectConfig) *Object { _ = "STUB: not implemented"; return nil }

func (gt *Object) ensureCache() { _ = "STUB: not implemented"; return }

func (gt *Object) AddFieldConfig(fieldName string, fieldConfig *Field) {
	_ = "STUB: not implemented"
	return
}

func (gt *Object) Name() string { _ = "STUB: not implemented"; return "" }

func (gt *Object) Description() string { _ = "STUB: not implemented"; return "" }

func (gt *Object) String() string { _ = "STUB: not implemented"; return "" }

func (gt *Object) Fields() FieldDefinitionMap {
	_ = "STUB: not implemented"
	return *new(FieldDefinitionMap)
}

func (gt *Object) Interfaces() []*Interface { _ = "STUB: not implemented"; return nil }

func (gt *Object) Error() error { _ = "STUB: not implemented"; return nil }

func defineInterfaces(ttype *Object, interfaces []*Interface) ([]*Interface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func defineFieldMap(ttype Named, fieldMap Fields) (FieldDefinitionMap, error) {
	_ = "STUB: not implemented"
	return *new(FieldDefinitionMap), nil
}

type ResolveParams struct {
	Source interface{}

	Args map[string]interface{}

	Info ResolveInfo

	Context context.Context
}

type FieldResolveFn func(p ResolveParams) (interface{}, error)

type ResolveInfo struct {
	FieldName      string
	FieldASTs      []*ast.Field
	Path           *ResponsePath
	ReturnType     Output
	ParentType     Composite
	Schema         Schema
	Fragments      map[string]ast.Definition
	RootValue      interface{}
	Operation      ast.Definition
	VariableValues map[string]interface{}
}

type Fields map[string]*Field

type Field struct {
	Name              string              `json:"name"`
	Type              Output              `json:"type"`
	Args              FieldConfigArgument `json:"args"`
	Resolve           FieldResolveFn      `json:"-"`
	Subscribe         FieldResolveFn      `json:"-"`
	DeprecationReason string              `json:"deprecationReason"`
	Description       string              `json:"description"`
}

type FieldConfigArgument map[string]*ArgumentConfig

type ArgumentConfig struct {
	Type         Input       `json:"type"`
	DefaultValue interface{} `json:"defaultValue"`
	Description  string      `json:"description"`
}

type FieldDefinitionMap map[string]*FieldDefinition
type FieldDefinition struct {
	Name              string         `json:"name"`
	Description       string         `json:"description"`
	Type              Output         `json:"type"`
	Args              []*Argument    `json:"args"`
	Resolve           FieldResolveFn `json:"-"`
	Subscribe         FieldResolveFn `json:"-"`
	DeprecationReason string         `json:"deprecationReason"`
}

type FieldArgument struct {
	Name         string      `json:"name"`
	Type         Type        `json:"type"`
	DefaultValue interface{} `json:"defaultValue"`
	Description  string      `json:"description"`
}

type Argument struct {
	PrivateName        string      `json:"name"`
	Type               Input       `json:"type"`
	DefaultValue       interface{} `json:"defaultValue"`
	PrivateDescription string      `json:"description"`
}

func (st *Argument) Name() string { _ = "STUB: not implemented"; return "" }

func (st *Argument) Description() string { _ = "STUB: not implemented"; return "" }

func (st *Argument) String() string { _ = "STUB: not implemented"; return "" }

func (st *Argument) Error() error { _ = "STUB: not implemented"; return nil }

type Interface struct {
	PrivateName        string `json:"name"`
	PrivateDescription string `json:"description"`
	ResolveType        ResolveTypeFn

	typeConfig        InterfaceConfig
	initialisedFields bool
	fields            FieldDefinitionMap
	err               error
}
type InterfaceConfig struct {
	Name        string      `json:"name"`
	Fields      interface{} `json:"fields"`
	ResolveType ResolveTypeFn
	Description string `json:"description"`
}

type ResolveTypeParams struct {
	Value interface{}

	Info ResolveInfo

	Context context.Context
}

type ResolveTypeFn func(p ResolveTypeParams) *Object

func NewInterface(config InterfaceConfig) *Interface { _ = "STUB: not implemented"; return nil }

func (it *Interface) AddFieldConfig(fieldName string, fieldConfig *Field) {
	_ = "STUB: not implemented"
	return
}

func (it *Interface) Name() string { _ = "STUB: not implemented"; return "" }

func (it *Interface) Description() string { _ = "STUB: not implemented"; return "" }

func (it *Interface) Fields() (fields FieldDefinitionMap) {
	_ = "STUB: not implemented"
	return *new(FieldDefinitionMap)
}

func (it *Interface) String() string { _ = "STUB: not implemented"; return "" }

func (it *Interface) Error() error { _ = "STUB: not implemented"; return nil }

type Union struct {
	PrivateName        string `json:"name"`
	PrivateDescription string `json:"description"`
	ResolveType        ResolveTypeFn

	typeConfig      UnionConfig
	initalizedTypes bool
	types           []*Object
	possibleTypes   map[string]bool

	err error
}

type UnionTypesThunk func() []*Object

type UnionConfig struct {
	Name        string      `json:"name"`
	Types       interface{} `json:"types"`
	ResolveType ResolveTypeFn
	Description string `json:"description"`
}

func NewUnion(config UnionConfig) *Union { _ = "STUB: not implemented"; return nil }

func (ut *Union) Types() []*Object { _ = "STUB: not implemented"; return nil }

func defineUnionTypes(objectType *Union, unionTypes []*Object) ([]*Object, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ut *Union) String() string { _ = "STUB: not implemented"; return "" }

func (ut *Union) Name() string { _ = "STUB: not implemented"; return "" }

func (ut *Union) Description() string { _ = "STUB: not implemented"; return "" }

func (ut *Union) Error() error { _ = "STUB: not implemented"; return nil }

type Enum struct {
	PrivateName        string `json:"name"`
	PrivateDescription string `json:"description"`

	enumConfig   EnumConfig
	values       []*EnumValueDefinition
	valuesLookup map[interface{}]*EnumValueDefinition
	nameLookup   map[string]*EnumValueDefinition

	err error
}
type EnumValueConfigMap map[string]*EnumValueConfig
type EnumValueConfig struct {
	Value             interface{} `json:"value"`
	DeprecationReason string      `json:"deprecationReason"`
	Description       string      `json:"description"`
}
type EnumConfig struct {
	Name        string             `json:"name"`
	Values      EnumValueConfigMap `json:"values"`
	Description string             `json:"description"`
}
type EnumValueDefinition struct {
	Name              string      `json:"name"`
	Value             interface{} `json:"value"`
	DeprecationReason string      `json:"deprecationReason"`
	Description       string      `json:"description"`
}

func NewEnum(config EnumConfig) *Enum { _ = "STUB: not implemented"; return nil }

func (gt *Enum) defineEnumValues(valueMap EnumValueConfigMap) ([]*EnumValueDefinition, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (gt *Enum) Values() []*EnumValueDefinition { _ = "STUB: not implemented"; return nil }

func (gt *Enum) Serialize(value interface{}) interface{} { _ = "STUB: not implemented"; return nil }

func (gt *Enum) ParseValue(value interface{}) interface{} { _ = "STUB: not implemented"; return nil }

func (gt *Enum) ParseLiteral(valueAST ast.Value) interface{} { _ = "STUB: not implemented"; return nil }

func (gt *Enum) Name() string { _ = "STUB: not implemented"; return "" }

func (gt *Enum) Description() string { _ = "STUB: not implemented"; return "" }

func (gt *Enum) String() string { _ = "STUB: not implemented"; return "" }

func (gt *Enum) Error() error { _ = "STUB: not implemented"; return nil }

func (gt *Enum) getValueLookup() map[interface{}]*EnumValueDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (gt *Enum) getNameLookup() map[string]*EnumValueDefinition {
	_ = "STUB: not implemented"
	return nil
}

type InputObject struct {
	PrivateName        string `json:"name"`
	PrivateDescription string `json:"description"`

	typeConfig InputObjectConfig
	fields     InputObjectFieldMap
	init       bool
	err        error
}
type InputObjectFieldConfig struct {
	Type         Input       `json:"type"`
	DefaultValue interface{} `json:"defaultValue"`
	Description  string      `json:"description"`
}
type InputObjectField struct {
	PrivateName        string      `json:"name"`
	Type               Input       `json:"type"`
	DefaultValue       interface{} `json:"defaultValue"`
	PrivateDescription string      `json:"description"`
}

func (st *InputObjectField) Name() string { _ = "STUB: not implemented"; return "" }

func (st *InputObjectField) Description() string { _ = "STUB: not implemented"; return "" }

func (st *InputObjectField) String() string { _ = "STUB: not implemented"; return "" }

func (st *InputObjectField) Error() error { _ = "STUB: not implemented"; return nil }

type InputObjectConfigFieldMap map[string]*InputObjectFieldConfig
type InputObjectFieldMap map[string]*InputObjectField
type InputObjectConfigFieldMapThunk func() InputObjectConfigFieldMap
type InputObjectConfig struct {
	Name        string      `json:"name"`
	Fields      interface{} `json:"fields"`
	Description string      `json:"description"`
}

func NewInputObject(config InputObjectConfig) *InputObject { _ = "STUB: not implemented"; return nil }

func (gt *InputObject) defineFieldMap() InputObjectFieldMap {
	_ = "STUB: not implemented"
	return *new(InputObjectFieldMap)
}

func (gt *InputObject) AddFieldConfig(fieldName string, fieldConfig *InputObjectFieldConfig) {
	_ = "STUB: not implemented"
	return
}

func (gt *InputObject) Fields() InputObjectFieldMap {
	_ = "STUB: not implemented"
	return *new(InputObjectFieldMap)
}

func (gt *InputObject) Name() string { _ = "STUB: not implemented"; return "" }

func (gt *InputObject) Description() string { _ = "STUB: not implemented"; return "" }

func (gt *InputObject) String() string { _ = "STUB: not implemented"; return "" }

func (gt *InputObject) Error() error { _ = "STUB: not implemented"; return nil }

type List struct {
	OfType Type `json:"ofType"`

	err error
}

func NewList(ofType Type) *List { _ = "STUB: not implemented"; return nil }

func (gl *List) Name() string { _ = "STUB: not implemented"; return "" }

func (gl *List) Description() string { _ = "STUB: not implemented"; return "" }

func (gl *List) String() string { _ = "STUB: not implemented"; return "" }

func (gl *List) Error() error { _ = "STUB: not implemented"; return nil }

type NonNull struct {
	OfType Type `json:"ofType"`

	err error
}

func NewNonNull(ofType Type) *NonNull { _ = "STUB: not implemented"; return nil }

func (gl *NonNull) Name() string { _ = "STUB: not implemented"; return "" }

func (gl *NonNull) Description() string { _ = "STUB: not implemented"; return "" }

func (gl *NonNull) String() string { _ = "STUB: not implemented"; return "" }

func (gl *NonNull) Error() error { _ = "STUB: not implemented"; return nil }

var NameRegExp = regexp.MustCompile("^[_a-zA-Z][_a-zA-Z0-9]*$")

func assertValidName(name string) error { _ = "STUB: not implemented"; return nil }

type ResponsePath struct {
	Prev *ResponsePath
	Key  interface{}
}

func (p *ResponsePath) WithKey(key interface{}) *ResponsePath {
	_ = "STUB: not implemented"
	return nil
}

func (p *ResponsePath) AsArray() []interface{} { _ = "STUB: not implemented"; return nil }
