package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 1) //buffer size 為1
	ch <- 1
	select {
	case ch <- 2 :
		fmt.Println("channel value is", <-ch)
		fmt.Println("channel value is", <-ch)
	default:
		fmt.Println("channel blocking")
	}
	
}
//如果將buffer size改為2，就可以繼續sent值進入channel