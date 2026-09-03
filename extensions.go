package graphql

import (
	"context"

	"github.com/graphql-go/graphql/gqlerrors"
)

type (
	ParseFinishFunc func(error)

	parseFinishFuncHandler func(error) []gqlerrors.FormattedError

	ValidationFinishFunc func([]gqlerrors.FormattedError)

	validationFinishFuncHandler func([]gqlerrors.FormattedError) []gqlerrors.FormattedError

	ExecutionFinishFunc func(*Result)

	executionFinishFuncHandler func(*Result) []gqlerrors.FormattedError

	ResolveFieldFinishFunc func(interface{}, error)

	resolveFieldFinishFuncHandler func(interface{}, error) []gqlerrors.FormattedError
)

type Extension interface {
	Init(context.Context, *Params) context.Context

	Name() string

	ParseDidStart(context.Context) (context.Context, ParseFinishFunc)

	ValidationDidStart(context.Context) (context.Context, ValidationFinishFunc)

	ExecutionDidStart(context.Context) (context.Context, ExecutionFinishFunc)

	ResolveFieldDidStart(context.Context, *ResolveInfo) (context.Context, ResolveFieldFinishFunc)

	HasResult() bool

	GetResult(context.Context) interface{}
}

func handleExtensionsInits(p *Params) gqlerrors.FormattedErrors {
	_ = "STUB: not implemented"
	return *new(gqlerrors.FormattedErrors)
}

func handleExtensionsParseDidStart(p *Params) ([]gqlerrors.FormattedError, parseFinishFuncHandler) {
	_ = "STUB: not implemented"
	return nil, *new(parseFinishFuncHandler)
}

func handleExtensionsValidationDidStart(p *Params) ([]gqlerrors.FormattedError, validationFinishFuncHandler) {
	_ = "STUB: not implemented"
	return nil, *new(validationFinishFuncHandler)
}

func handleExtensionsExecutionDidStart(p *ExecuteParams) ([]gqlerrors.FormattedError, executionFinishFuncHandler) {
	_ = "STUB: not implemented"
	return nil, *new(executionFinishFuncHandler)
}

func handleExtensionsResolveFieldDidStart(exts []Extension, p *executionContext, i *ResolveInfo) ([]gqlerrors.FormattedError, resolveFieldFinishFuncHandler) {
	_ = "STUB: not implemented"
	return nil, *new(resolveFieldFinishFuncHandler)
}

func addExtensionResults(p *ExecuteParams, result *Result) { _ = "STUB: not implemented"; return }
