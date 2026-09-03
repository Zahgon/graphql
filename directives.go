package graphql

const (
	DirectiveLocationQuery              = "QUERY"
	DirectiveLocationMutation           = "MUTATION"
	DirectiveLocationSubscription       = "SUBSCRIPTION"
	DirectiveLocationField              = "FIELD"
	DirectiveLocationFragmentDefinition = "FRAGMENT_DEFINITION"
	DirectiveLocationFragmentSpread     = "FRAGMENT_SPREAD"
	DirectiveLocationInlineFragment     = "INLINE_FRAGMENT"

	DirectiveLocationSchema               = "SCHEMA"
	DirectiveLocationScalar               = "SCALAR"
	DirectiveLocationObject               = "OBJECT"
	DirectiveLocationFieldDefinition      = "FIELD_DEFINITION"
	DirectiveLocationArgumentDefinition   = "ARGUMENT_DEFINITION"
	DirectiveLocationInterface            = "INTERFACE"
	DirectiveLocationUnion                = "UNION"
	DirectiveLocationEnum                 = "ENUM"
	DirectiveLocationEnumValue            = "ENUM_VALUE"
	DirectiveLocationInputObject          = "INPUT_OBJECT"
	DirectiveLocationInputFieldDefinition = "INPUT_FIELD_DEFINITION"
)

const DefaultDeprecationReason = "No longer supported"

var SpecifiedDirectives = []*Directive{
	IncludeDirective,
	SkipDirective,
	DeprecatedDirective,
}

type Directive struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Locations   []string    `json:"locations"`
	Args        []*Argument `json:"args"`

	err error
}

type DirectiveConfig struct {
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Locations   []string            `json:"locations"`
	Args        FieldConfigArgument `json:"args"`
}

func NewDirective(config DirectiveConfig) *Directive { _ = "STUB: not implemented"; return nil }

var IncludeDirective = NewDirective(DirectiveConfig{
	Name: "include",
	Description: "Directs the executor to include this field or fragment only when " +
		"the `if` argument is true.",
	Locations: []string{
		DirectiveLocationField,
		DirectiveLocationFragmentSpread,
		DirectiveLocationInlineFragment,
	},
	Args: FieldConfigArgument{
		"if": &ArgumentConfig{
			Type:        NewNonNull(Boolean),
			Description: "Included when true.",
		},
	},
})

var SkipDirective = NewDirective(DirectiveConfig{
	Name: "skip",
	Description: "Directs the executor to skip this field or fragment when the `if` " +
		"argument is true.",
	Args: FieldConfigArgument{
		"if": &ArgumentConfig{
			Type:        NewNonNull(Boolean),
			Description: "Skipped when true.",
		},
	},
	Locations: []string{
		DirectiveLocationField,
		DirectiveLocationFragmentSpread,
		DirectiveLocationInlineFragment,
	},
})

var DeprecatedDirective = NewDirective(DirectiveConfig{
	Name:        "deprecated",
	Description: "Marks an element of a GraphQL schema as no longer supported.",
	Args: FieldConfigArgument{
		"reason": &ArgumentConfig{
			Type: String,
			Description: "Explains why this element was deprecated, usually also including a " +
				"suggestion for how to access supported similar data. Formatted" +
				"in [Markdown](https://daringfireball.net/projects/markdown/).",
			DefaultValue: DefaultDeprecationReason,
		},
	},
	Locations: []string{
		DirectiveLocationFieldDefinition,
		DirectiveLocationEnumValue,
	},
})
