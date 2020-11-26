package main

import (
	"fmt"
)

//Point is a type
type Point struct{	//將type宣告所使用的struct為class
	x float64		//將 x 和 y 改為小寫，代表該屬性是私有屬性
	y float64
}

func newpoint(x, y float64) *Point  { //建構的func作為obiect
	p := new(Point)
	p.SetX(x) 
	p.SetY(y)
	//該類別的 setters 來初始化屬性，這是刻意的動作。因為我們要確保在設置屬性時的行為保持一致
	return p
}

func (p *Point) X() float64  {
	return p.x
}

func (p *Point) Y() float64 {
	return p.y
}
func (p *Point) SetX(x float64){
	p.x = x
}
func (p *Point) SetY(y float64){
	p.y = y
}

func main()  {
	p := newpoint(0, 0)
	fmt.Println(p.X(), p.Y())
	p.SetX(3)
	p.SetY(4)
	fmt.Println(p.X(), p.Y())

}