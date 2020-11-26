package main

import (
	"fmt"
)

type Integer int

func (x Integer) Equal(y Integer) bool  {
	return x == y
}
func (x Integer) LessThan(y Integer) bool  {
	return x < y
}
func (x Integer) MoreThan(y Integer) bool {
	return x > y
}
func (x *Integer) Increase(y Integer) {
	*x = *x + y
}
func (x *Integer) Decrease(y Integer)  {
	*x = *x - y
}

type IntNumber interface{
	Equal(y Integer) bool
	LessThan(y Integer) bool
	MoreThan(y Integer) bool
	Increase(y Integer)
	Decrease(y Integer)
}

//將 type Integer 的值賦予給 InNumber
func main()  {
	var a Integer = 1
//由於 method Increase(),Descrease()為point method，若Integer未傳遞point，則無法實現所有method
	var b IntNumber = &a
	fmt.Println(b.Equal(1))
}