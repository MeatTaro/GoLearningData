package main

import (
	"fmt"
)
//利用函式呼叫函式，計算費伯納係數
func factorial (x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}
func main()  {
	fmt.Println(factorial(3))
}