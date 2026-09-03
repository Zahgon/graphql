package ast

type Value interface {
	GetValue() interface{}
	GetKind() string
	GetLoc() *Location
}

var _ Value = (*Variable)(nil)
var _ Value = (*IntValue)(nil)
var _ Value = (*FloatValue)(nil)
var _ Value = (*StringValue)(nil)
var _ Value = (*BooleanValue)(nil)
var _ Value = (*EnumValue)(nil)
var _ Value = (*ListValue)(nil)
var _ Value = (*ObjectValue)(nil)

type Variable struct {
	Kind string
	Loc  *Location
	Name *Name
}

func NewVariable(v *Variable) *Variable { _ = "STUB: not implemented"; return nil }

func (v *Variable) GetKind() string { _ = "STUB: not implemented"; return "" }

func (v *Variable) GetLoc() *Location { _ = "STUB: not implemented"; return nil }

func (v *Variable) GetValue() interface{} { _ = "STUB: not implemented"; return nil }

func (v *Variable) GetName() interface{} { _ = "STUB: not implemented"; return nil }

type IntValue struct {
	Kind  string
	Loc   *Location
	Value string
}

func NewIntValue(v *IntValue) *IntValue { _ = "STUB: not implemented"; return nil }

func (v *IntValue) GetKind() string { _ = "STUB: not implemented"; return "" }

func (v *IntValue) GetLoc() *Location { _ = "STUB: not implemented"; return nil }

func (v *IntValue) GetValue() interface{} { _ = "STUB: not implemented"; return nil }

type FloatValue struct {
	Kind  string
	Loc   *Location
	Value string
}

func NewFloatValue(v *FloatValue) *FloatValue { _ = "STUB: not implemented"; return nil }

func (v *FloatValue) GetKind() string { _ = "STUB: not implemented"; return "" }

func (v *FloatValue) GetLoc() *Location { _ = "STUB: not implemented"; return nil }

func (v *FloatValue) GetValue() interface{} { _ = "STUB: not implemented"; return nil }

type StringValue struct {
	Kind  string
	Loc   *Location
	Value string
}

func NewStringValue(v *StringValue) *StringValue { _ = "STUB: not implemented"; return nil }

func (v *StringValue) GetKind() string { _ = "STUB: not implemented"; return "" }

func (v *StringValue) GetLoc() *Location { _ = "STUB: not implemented"; return nil }

func (v *StringValue) GetValue() interface{} { _ = "STUB: not implemented"; return nil }

type BooleanValue struct {
	Kind  string
	Loc   *Location
	Value bool
}

func NewBooleanValue(v *BooleanValue) *BooleanValue { _ = "STUB: not implemented"; return nil }

func (v *BooleanValue) GetKind() string { _ = "STUB: not implemented"; return "" }

func (v *BooleanValue) GetLoc() *Location { _ = "STUB: not implemented"; return nil }

func (v *BooleanValue) GetValue() interface{} { _ = "STUB: not implemented"; return nil }

type EnumValue struct {
	Kind  string
	Loc   *Location
	Value string
}

func NewEnumValue(v *EnumValue) *EnumValue { _ = "STUB: not implemented"; return nil }

func (v *EnumValue) GetKind() string { _ = "STUB: not implemented"; return "" }

func (v *EnumValue) GetLoc() *Location { _ = "STUB: not implemented"; return nil }

func (v *EnumValue) GetValue() interface{} { _ = "STUB: not implemented"; return nil }

type ListValue struct {
	Kind   string
	Loc    *Location
	Values []Value
}

func NewListValue(v *ListValue) *ListValue { _ = "STUB: not implemented"; return nil }

func (v *ListValue) GetKind() string { _ = "STUB: not implemented"; return "" }

func (v *ListValue) GetLoc() *Location { _ = "STUB: not implemented"; return nil }

func (v *ListValue) GetValue() interface{} { _ = "STUB: not implemented"; return nil }

func (v *ListValue) GetValues() interface{} { _ = "STUB: not implemented"; return nil }

type ObjectValue struct {
	Kind   string
	Loc    *Location
	Fields []*ObjectField
}

func NewObjectValue(v *ObjectValue) *ObjectValue { _ = "STUB: not implemented"; return nil }

func (v *ObjectValue) GetKind() string { _ = "STUB: not implemented"; return "" }

func (v *ObjectValue) GetLoc() *Location { _ = "STUB: not implemented"; return nil }

func (v *ObjectValue) GetValue() interface{} { _ = "STUB: not implemented"; return nil }

type ObjectField struct {
	Kind  string
	Name  *Name
	Loc   *Location
	Value Value
}

func NewObjectField(f *ObjectField) *ObjectField { _ = "STUB: not implemented"; return nil }

func (f *ObjectField) GetKind() string { _ = "STUB: not implemented"; return "" }

func (f *ObjectField) GetLoc() *Location { _ = "STUB: not implemented"; return nil }

func (f *ObjectField) GetValue() interface{} { _ = "STUB: not implemented"; return nil }
