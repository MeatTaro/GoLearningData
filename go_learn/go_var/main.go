package main

import (
	"fmt"
)

var a string
//---------------
var (
	c string
	d int
)
//---------------
var e string = "3. 變數同時宣告並賦值"
//-----------------


func main()  {
	a = "1. 宣告一個變數並賦值"
	fmt.Println(a)
	var b int //變數宣告可於func外或者func內
	b = 1
	fmt.Println(b)
//---------------
	c = "2. 一次宣告多筆變數並賦值 "
	d = 0
	fmt.Println(c)
	fmt.Println(d)
//---------------
	fmt.Println(e)  
	var f int = 101 
	fmt.Println(f) 
//---------------
	g := "變數、同時宣告並賦值 簡寫"
	fmt.Println(g)
//---------------
	
	const(
		monday = iota + 1
		tuesday
		wednsday
	)
	fmt.Println(monday)
	fmt.Println(tuesday)
	fmt.Println(wednsday)
	//const為常數不可修改
}

//注意! 於func外宣告變數時為"全域變數"，若不希望變數互相干擾，盡量宣告於func內