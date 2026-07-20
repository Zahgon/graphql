package gqlerrors

import (
	"github.com/graphql-go/graphql/language/location"
)

type ExtendedError interface {
	error
	Extensions() map[string]interface{}
}

type FormattedError struct {
	Message       string                    `json:"message"`
	Locations     []location.SourceLocation `json:"locations"`
	Path          []interface{}             `json:"path,omitempty"`
	Extensions    map[string]interface{}    `json:"extensions,omitempty"`
	originalError error
}

func (g FormattedError) OriginalError() error { _ = "STUB: not implemented"; return nil }

func (g FormattedError) Error() string { _ = "STUB: not implemented"; return "" }

func NewFormattedError(message string) FormattedError {
	_ = "STUB: not implemented"
	return *new(FormattedError)
}

func FormatError(err error) FormattedError { _ = "STUB: not implemented"; return *new(FormattedError) }

func FormatErrors(errs ...error) []FormattedError { _ = "STUB: not implemented"; return nil }
