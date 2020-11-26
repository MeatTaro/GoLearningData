package main

import (
	"fmt"
)
//1. package 由func構成
//2. func 函式名稱(定義傳入值 傳入值型態) 函式回傳型態{函式內容}

func hello(name string){
	fmt.Printf("Hello %s", name)
}
func main()  {
	hello("MeatTaro")
}