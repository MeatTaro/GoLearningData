package main

import (
	"fmt"
)

func test(ch chan <- int) {
	for i:=0; i<5; i++ {	
		ch <- i
	}
	close(ch)
}

func main()  {
	ch := make(chan int, 5)
	go test(ch)
	for i := range ch {
		fmt.Println(i)
	}
}