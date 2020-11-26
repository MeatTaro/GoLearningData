package main

import (
	"fmt"
)

func main()  {
	ch := make(chan int)
	go func() {
		for i:=0; i<5; i++ {
			fmt.Printf("Sender: sanding data %v \n",i)
			ch <- i
		}
		fmt.Println("Sender: close the channel")
		close(ch)
	}()

	for {
		num, ok := <- ch
		if !ok {
			fmt.Println("Receiver: channel closed")
			break
		}
		fmt.Printf("Receiver: receive data %v \n", num)
	}
	fmt.Println("end")
}