package benchutil

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

type color struct {
	Hex string
	R   int
	G   int
	B   int
}

func ListSchemaWithXItems(x int) graphql.Schema {
	_ = "STUB: not implemented"
	return *new(graphql.Schema)
}

var colors []color

func init() {
	colors = make([]color, 0, 256*16*16)

	for r := 0; r < 256; r++ {
		for g := 0; g < 16; g++ {
			for b := 0; b < 16; b++ {
				colors = append(colors, color{
					Hex: fmt.Sprintf("#%x%x%x", r, g, b),
					R:   r,
					G:   g,
					B:   b,
				})
			}
		}
	}
}

func generateXListItems(x int) []color { _ = "STUB: not implemented"; return nil }
