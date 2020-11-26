package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	timeOut := make(chan bool)

	go func() {
		time.Sleep(1e9)
		timeOut <- true
	}()

	select {
		case <- ch :
			fmt.Println(ch)
		case <- timeOut :
			fmt.Println("time out")
	}
}