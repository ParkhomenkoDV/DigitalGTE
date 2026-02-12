package cc

import "github.com/ParkhomenkoDV/DigitalGTE.git/substance"

type CC struct {
}

func New() *CC {
	return &CC{}
}

func (cc *CC) Predict() {

}
func (cc *CC) Equations() {

}
func (cc *CC) Calculate() substance.Substance {
	return substance.Substance{}
}

func (cc *CC) Solve() substance.Substance {
	return substance.Substance{}
}
