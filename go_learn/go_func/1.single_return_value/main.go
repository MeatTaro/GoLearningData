package main

import "fmt"

func add(i, j int) int {
	return i + j //單一回傳值
}

func main() {
	fmt.Println(add(1, 2))
}