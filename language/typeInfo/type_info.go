package typeInfo

import (
	"github.com/graphql-go/graphql/language/ast"
)

type TypeInfoI interface {
	Enter(node ast.Node)
	Leave(node ast.Node)
}
