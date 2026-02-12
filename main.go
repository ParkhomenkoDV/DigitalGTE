package main

import (
	"github.com/ParkhomenkoDV/DigitalGTE.git/gte"
	"github.com/ParkhomenkoDV/DigitalGTE.git/gte/cc"
	"github.com/ParkhomenkoDV/DigitalGTE.git/gte/node"
	"github.com/ParkhomenkoDV/DigitalGTE.git/gte/schema"
	"github.com/ParkhomenkoDV/DigitalGTE.git/gte/turbocompressor/compressor"
	"github.com/ParkhomenkoDV/DigitalGTE.git/gte/turbocompressor/turbine"
)

func main() {
	compressor := compressor.New()
	cc := cc.New()
	turbine := turbine.New()

	_ = compressor
	_ = cc
	_ = turbine

	schema := schema.New(
		map[int][]node.Node{
			1: []node.Node{compressor, cc, turbine},
		},
	)

	gte := gte.New(*schema)

	gte.Solve()
}
