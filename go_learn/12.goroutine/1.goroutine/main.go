package main

import (
	// "time"
	// "math/rand"
	"fmt"
)

func f(n int) {
	for i:=0;i<10;i++{
		fmt.Println(n, ":" ,i)
		// amt := time.Duration(rand.Intn(250))
		// time.Sleep(time.Millisecond * amt)
	}
}

func main()  {
	for i:=0;i<10;i++{
		go f(i)
	}
	go f(100)
	var input string
	fmt.Scanln(&input)
}