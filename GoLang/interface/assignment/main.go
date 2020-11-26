package main

import (
	"fmt"
	"assignment/oop2"
	"assignment/oop1"
)
//透過interface進行assignment操作
//1.if a type implement an interface, then the feilds of the type could be assigned to the interface. 
//(將實現接口的方法的'值'賦值給接口)

type Number int

func (x Number) Equal(y int) bool  {
	return int(x) == y
}
func (x Number) LessThan(y int) bool {
	return int(x) < y
}
func (x Number) MoreThan(y int) bool  {
	return int(x) > y
}
func (x *Number) Increase(y int) {
	*x = *x + Number(y)
}
func (x *Number) Decrease(y int) {
	*x = *x - Number(y)
}

func main()  {
	var num1 Number =1
	//2.assign feild to interface with point. because point could implement all method in the interface.(include Increse() Decrease()),unless the method in type didn't use any point.
	var num2 oop2.Number2 = &num1
	//3. an interface could be assigned to another interface
	//any type that implement the interface 'Number1',it implement 'Number2' too.
	//any feild in method implement the interface could assign to 'Number2' 
	var num3 oop1.Number1 = num2
	fmt.Println(num2.Equal(1))
	fmt.Println(num3.Equal(2))
	num2.Increase(3)
	fmt.Println(num1)
}