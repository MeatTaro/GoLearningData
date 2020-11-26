package main

import (
	"fmt"
)

func main()  {
	chanString := make(chan string)
	go func() {
		chanString <- "Hello World"
	}()
	fmt.Println(<- chanString)
}