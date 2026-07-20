package graphql

import (
	"context"
)

type Params struct {
	Schema Schema

	RequestString string

	RootObject map[string]interface{}

	VariableValues map[string]interface{}

	OperationName string

	Context context.Context
}

func Do(p Params) *Result { _ = "STUB: not implemented"; return nil }
