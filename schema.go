package graphql

type SchemaConfig struct {
	Query        *Object
	Mutation     *Object
	Subscription *Object
	Types        []Type
	Directives   []*Directive
	Extensions   []Extension
}

type TypeMap map[string]Type

type Schema struct {
	typeMap    TypeMap
	directives []*Directive

	queryType        *Object
	mutationType     *Object
	subscriptionType *Object
	implementations  map[string][]*Object
	possibleTypeMap  map[string]map[string]bool
	extensions       []Extension
}

func NewSchema(config SchemaConfig) (Schema, error) {
	_ = "STUB: not implemented"
	return *new(Schema), nil
}

func (gq *Schema) AddImplementation() error { _ = "STUB: not implemented"; return nil }

func (gq *Schema) AppendType(objectType Type) error { _ = "STUB: not implemented"; return nil }

func (gq *Schema) QueryType() *Object { _ = "STUB: not implemented"; return nil }

func (gq *Schema) MutationType() *Object { _ = "STUB: not implemented"; return nil }

func (gq *Schema) SubscriptionType() *Object { _ = "STUB: not implemented"; return nil }

func (gq *Schema) Directives() []*Directive { _ = "STUB: not implemented"; return nil }

func (gq *Schema) Directive(name string) *Directive { _ = "STUB: not implemented"; return nil }

func (gq *Schema) TypeMap() TypeMap { _ = "STUB: not implemented"; return *new(TypeMap) }

func (gq *Schema) Type(name string) Type { _ = "STUB: not implemented"; return *new(Type) }

func (gq *Schema) PossibleTypes(abstractType Abstract) []*Object {
	_ = "STUB: not implemented"
	return nil
}

func (gq *Schema) IsPossibleType(abstractType Abstract, possibleType *Object) bool {
	_ = "STUB: not implemented"
	return false
}

func (gq *Schema) AddExtensions(e ...Extension) { _ = "STUB: not implemented"; return }

func typeMapReducer(schema *Schema, typeMap TypeMap, objectType Type) (TypeMap, error) {
	_ = "STUB: not implemented"
	return *new(TypeMap), nil
}

func assertObjectImplementsInterface(schema *Schema, object *Object, iface *Interface) error {
	_ = "STUB: not implemented"
	return nil
}

func isEqualType(typeA Type, typeB Type) bool { _ = "STUB: not implemented"; return false }

func isTypeSubTypeOf(schema *Schema, maybeSubType Type, superType Type) bool {
	_ = "STUB: not implemented"
	return false
}
