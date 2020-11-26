package main

import (
	"fmt"
)

type Point struct {
	x,y float64
}
func newPoint(x,y float64) *Point  {
	p := new(Point)
	p.SetX(x)
	p.SetY(y)
	return p
}
func (p *Point) SetX(x float64)  {
	p.x = x
}
func (p *Point) SetY(y float64)  {
	p.y = y
}
func (p *Point) X() float64  {
	return p.x
}
func (p *Point) Y() float64  {
	return p.y
}
type Point3D struct {
	Point
	z float64
}
func newPoint3D(x,y,z float64) *Point3D {
	p := new(Point3D)
	p.SetX(x)
	p.SetY(y)
	p.SetZ(z)
	return p
}
func (p *Point3D) SetZ(z float64)  {
	p.z = z
}
func (p *Point3D) Z() float64  {
	return p.z
}
func main()  {
	p := newPoint3D(0,0,0)
	fmt.Println(p.X(),p.Y(),p.Z())
}

