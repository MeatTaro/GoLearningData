package main

import (
	"fmt"
)

func main()  {
	defer func () {
		fmt.Println("clean")
	}()
	var i = 1
	var j = 0
	if j == 0 {
		panic("fuck")
	}
	k := i/j
	fmt.Printf("%d / %d = %d \n ",i,j,k)
}