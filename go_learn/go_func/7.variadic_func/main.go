package main

import (
	"fmt"
)

func add(x ... int)int { //x ... 可表示零或零個以上的參數
	total := 0
	for _,v:= range x{
		total += v
	}
	return total
}
func main()  {
	fmt.Println(add(1,2,3)) //add(1,2,3) 表示傳入此三個參數進入add(x... int)
	xs:=[]int{1,2,3}
	fmt.Println(add(xs...)) //也可以用於傳遞slice
}