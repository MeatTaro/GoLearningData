package main

import "fmt"

func swap(i, j int) (int, int) { //(int, int) 定義回傳函式型態為2個數值
	return j, i //多重回傳值 (本案例回傳值進行交換)
}

func main() {
	a := 1
	b := 2
	a, b = swap(1, 2)
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	a, b = b, a
	fmt.Println("a:", a)
	fmt.Println("b:", b)
}