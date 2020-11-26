package main

import (
	"math"
	"fmt"
)
//Point is a type
type Point struct{
	x,y float64
}
//Circle is a type
type Circle struct{
	r float64
	p Point
}

func circleArea(c *Circle) float64  {
	return math.Pi * c.r*c.r
}
func main()  {
	// c := new(Circle)
	// c.p.x = 10
	// c.p.y = 5
	// c.r = 10
	c := Circle{10.0, Point{x:0,y:0}}
	fmt.Println("中心座標:",c.p.x,c.p.y)
	fmt.Println("圓半徑:",c.r)
	fmt.Println("面積:",circleArea(&c))
}