package main

import (
	"fmt"
	"time"
)

func main() {
	timeout := make(chan bool,1)
	go func()  {
		time.Sleep(5*time.Second)
		timeout <- true
	}()
	ch := make(chan int)

	select {
		case <- ch :
		case <- timeout :
			fmt.Println("time out")
	}
}
