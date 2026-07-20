package testutil

import (
	"testing"

	"github.com/graphql-go/graphql"
)

type TestResponse struct {
	Data   string
	Errors []string
}

type TestSubscription struct {
	Name            string
	Schema          graphql.Schema
	Query           string
	OperationName   string
	Variables       map[string]interface{}
	ExpectedResults []TestResponse
}

func RunSubscribes(t *testing.T, tests []*TestSubscription) { _ = "STUB: not implemented"; return }

func RunSubscribe(t *testing.T, test *TestSubscription) { _ = "STUB: not implemented"; return }

func checkErrorStrings(t *testing.T, expected, actual []string) { _ = "STUB: not implemented"; return }

func formatJSON(data string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func pretty(x interface{}) string { _ = "STUB: not implemented"; return "" }
