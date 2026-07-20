package graphql

type SubscribeParams struct {
	Schema        Schema
	RequestString string
	RootValue     interface{}

	VariableValues  map[string]interface{}
	OperationName   string
	FieldResolver   FieldResolveFn
	FieldSubscriber FieldResolveFn
}

func Subscribe(p Params) chan *Result { _ = "STUB: not implemented"; return nil }

func sendOneResultAndClose(res *Result) chan *Result { _ = "STUB: not implemented"; return nil }

func ExecuteSubscription(p ExecuteParams) chan *Result { _ = "STUB: not implemented"; return nil }
