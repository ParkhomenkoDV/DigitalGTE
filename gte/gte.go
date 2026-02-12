package gte

import "github.com/ParkhomenkoDV/DigitalGTE.git/gte/schema"

type GTE struct {
	schema.Schema
}

func New(schema schema.Schema) *GTE {
	return &GTE{
		schema,
	}
}

func (gte *GTE) Solve() {

}
