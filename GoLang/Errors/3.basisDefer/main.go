package main

import (
	"fmt"
)

func printError()  {
	fmt.Println("XXXXX")
}
func main() {
	defer printError()
	defer func() {
		fmt.Println("XDDDDDD")
	}()

	var i=1
	var j=0
	var k=i/j

	fmt.Printf("%d / %d = %d \n ",i,j,k)
}