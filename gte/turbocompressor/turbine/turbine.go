package turbine

import "github.com/ParkhomenkoDV/DigitalGTE.git/substance"

type Turbine struct {
}

func New() *Turbine {
	return &Turbine{}
}

func (t *Turbine) Predict() {

}
func (t *Turbine) Equations() {

}
func (t *Turbine) Calculate() substance.Substance {
	return substance.Substance{}
}

func (t *Turbine) Solve() substance.Substance {
	return substance.Substance{}
}
