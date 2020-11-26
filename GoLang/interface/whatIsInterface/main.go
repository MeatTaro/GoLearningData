package main

import (
	"fmt"
	"math"
)

//Interface is another type, using to named collections of method signatures.

type geometry interface{
	area() float64
	perim() float64
}

//values of interface type can hold any value that implements those methds'area()' 'perim()'

type Rect struct {
	width float64
	height float64
}
type Circle struct{
	radious float64
}

func (x Rect) area() float64 {
	return x.width * x.height
}

func (x Circle) area() float64 {
	return x.radious * x.radious * math.Pi
}

func (x Rect) perim() float64 {
	return x.width * 2 + x.height * 2 
}

func (x Circle) perim() float64 {
	return x.radious * 2 * math.Pi
}

func mesure(x geometry)  {
	fmt.Println(x.area())
	fmt.Println(x.perim())
}

func main()  {
	r := Rect{width:1, height:1}
	c := Circle{radious:1}
	mesure(r)
	mesure(c)
}