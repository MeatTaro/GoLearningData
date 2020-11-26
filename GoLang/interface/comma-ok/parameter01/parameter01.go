package parameter01

import (
	"math"
)

type Parameter01 interface{
	Stright(i int) float64
	Circle(i int) float64
}

type Parameter float64

func (x Parameter) Stright(i int) float64   {
	return float64(x) * float64(i)
}

func (x Parameter) Circle(i int) float64 {
	return math.Pi * float64(x) * 2 * float64(i)
}

func (x Parameter) Hole(i int) float64 {
	return math.Pi * float64(x) * float64(i) * float64(i)
}
