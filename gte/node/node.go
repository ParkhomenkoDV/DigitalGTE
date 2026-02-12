package node

import "github.com/ParkhomenkoDV/DigitalGTE.git/substance"

type Node interface {
	Predict()
	Equations()
	Calculate() substance.Substance
	Solve() substance.Substance
}
