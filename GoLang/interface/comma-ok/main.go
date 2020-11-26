package main

import (
	"reflect"
	"comma-ok/process"
	"fmt"
	"comma-ok/parameter02"
	"comma-ok/parameter01"
)
//comma-ok(assertions)的語法是: value, ok := element.(T) ，element is var of interface，T is normal type， 如果comma-ok失敗 ok為false，comma-ok成功 ok為true，and value is assigned to the var
//用於判斷並轉化type的語法

//comma-ok-interface

func main()  {
	var qty1 parameter01.Parameter = 1
	var qty2 parameter02.Parameter02 = qty1
	// comma-ok if the var 'qty2' assigned to the interface 'Parameter02', then it can be assigned to Parameter01 too or not.
	//此語法qty2.(parameter01.Parameter01)不僅判斷，若 ok = true還進行 var'qty2' assign to 'qty3' 的操作，並執行if{}中的code
	if qty3, ok := qty2.(parameter01.Parameter01); ok {
		fmt.Println(qty3.Stright(5))
		fmt.Println(reflect.TypeOf(qty3))
	}

	var state = process.Station{"車床"}
	var iprocess process.Iprocess = process.CNC{state}
	//the value of state, that hold from type CNC assigned to the var 'iprocess' that type is interface'Iprocess'
	//then comma-ok if the iprocess's type could be assigned to CNC
	if cnc, ok := iprocess.(process.CNC); ok {
		fmt.Println(cnc.GetName())
		fmt.Println(reflect.TypeOf(cnc))
	}

}