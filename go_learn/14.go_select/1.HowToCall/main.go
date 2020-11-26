package main

import (
	"fmt"
)

func main()  {
	ch := make(chan int, 1)
	ch <- 100
	select {
		case <-ch:
			fmt.Println("random 01")
		case <-ch:
			fmt.Println("random 02")
		case <-ch:
			fmt.Println("random 03")
		case <-ch:
			fmt.Println("random 04")
		case <-ch:
			fmt.Println("random 05")
		case <-ch:
			fmt.Println("random 06")
		default:
			fmt.Println("exit")
	}
}