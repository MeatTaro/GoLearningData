package main

import (
	"math"
	"fmt"
)
//Point is a type 
//建立一個型別(type)名為Point的結構(struct)
//由結構(struct)組成
//struct內含有x,y二屬性(field)，其型別皆為float64
type Point struct {
    x float64
    y float64
}
//建構一函式(func)作為物件(object)
//函式傳入座標x,y
//用newc函式初始化
//利用私有函式定義x,y的值
//回傳Point屬性
func NewPoint(x,y float64) *Point {
	p := new(Point)
	p.SetX(x)
	p.SetY(y)
	return p
}
func (p *Point) SetX(x float64) {
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
type Point3D struct{
	Point
	z float64
}
func NewPoint3D(x,y,z float64) *Point3D {
	p := new(Point3D)
	p.SetX(x)
	p.SetY(y)
	p.SetZ(z)
	return p
}
func (p *Point3D) SetZ(z float64) {
	p.z = z
}
func (p *Point3D) Z() float64 {
	return p.z
}
//Interface 與 struct 一樣，都是使用 type 關鍵字建立的，後面跟著一個名稱，以及關鍵字 interface。不過這裡不定義欄位，而是定義一組 method，一組 method 指的是一串 methods，為了要實做 interface，所以這些 methods 都要有一個型別。
type IPoint interface{
	X() float64
	Y() float64
}
//p1、p2分別為Point、Point3D兩個各自獨立的類型
//因此新增一個interface使X()、Y()可以繼承IPint實作
func dist(p1,p2 IPoint) float64 {
	xSqr := p1.X() - p2.X()
	ySqr := p1.Y() - p2.Y()
	return math.Sqrt(xSqr*xSqr + ySqr*ySqr)
}
func main()  {
	p := new(Point)
	fmt.Println(p.x,p.y)
	p1 := NewPoint(1,1)
	fmt.Println(p1.X(), p1.Y())
	p2 := NewPoint3D(2,2,2)
	fmt.Println(p2.X(),p2.Y(),p2.Z())
	fmt.Println(dist(p1,p2))
}
