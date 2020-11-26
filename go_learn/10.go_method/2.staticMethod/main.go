package main

import (
	"fmt"
	"math"
)

//靜態方法 (Static Method)

type Point struct{
	x,y float64
}

func newpoint(x,y float64) *Point  {
	p := new(Point)
	p.SetX(x)
	p.SetY(y)
	return p
}
//SetX is method
func (p *Point) SetX(x float64) {
	p.x = x
}
//SetY is method
func (p *Point) SetY(y float64) {
	p.y = y
}
//X is method
func (p *Point) X() float64 {
	return p.x
}
//Y is method
func (p *Point) Y() float64 {
	return p.y
}
func dist(p1, p2 *Point) float64 {
	xSqr := p1.X() - p2.X()
	ySqr := p1.Y() - p2.Y()
	return math.Sqrt(xSqr*xSqr + ySqr*ySqr)
}
func main()  {
	p1 := newpoint(0,0)
	p2 := newpoint(3,4)
	fmt.Println(dist(p1,p2))
}