package combustionChamber

import (
	"github.com/ParkhomenkoDV/DigitalGTE.git/substance"
)

type Variables struct {
	EfficiencyBurn     float64
	PressureEfficiency float64
}

type CombustionChamber struct {
	Name string
	Characteristic
	Variables
}

func New(name string, characteristic Characteristic) *CombustionChamber {
	return &CombustionChamber{
		Name:           name,
		Characteristic: characteristic,
	}
}

func (cc *CombustionChamber) Figure() [2][]float64 {
	return [2][]float64{}
}

func (cc *CombustionChamber) Predict() {

}

func (cc *CombustionChamber) Equations() {

}

func (cc *CombustionChamber) Calculate(inlets ...*substance.Substance) *substance.Substance {
	return &substance.Substance{}
}

func (cc *CombustionChamber) Solve(inlets ...*substance.Substance) ([2]*substance.Substance, error) {
	outlet := substance.New(
		"exhaust",
		substance.Parameters{},
		substance.Functions{},
	)
	return [2]*substance.Substance{outlet}, nil
}
