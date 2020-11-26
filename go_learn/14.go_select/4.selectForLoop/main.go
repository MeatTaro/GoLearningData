package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	ch := make(chan string, 0) //buffer size is 0
	
	defer func() {
		close(ch)
	}()

	go func() {

		fmt.Println("start")
	LOOP:
		for {
			time.Sleep(1*time.Second)
			fmt.Println(time.Now().Unix())
			i++

			select {
			case m := <- ch : 
				fmt.Println(m)
				break LOOP
			default:
				// fmt.Println("blocking")
			}
		}
	}()

	time.Sleep(4*time.Second)
	ch <- "stop"
}