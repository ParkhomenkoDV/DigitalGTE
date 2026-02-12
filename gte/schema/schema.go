package schema

import "github.com/ParkhomenkoDV/DigitalGTE.git/gte/node"

type Schema struct {
	schema map[int][]node.Node
}

func New(schema map[int][]node.Node) *Schema {
	return &Schema{
		schema: schema,
	}
}
