package main

import (
	"time"
	"fmt"
)

func add(x, y int) {
	var z=x+y
	fmt.Printf("%d + %d = %d \n", x,y,z)
}

func main()  {
	
	for i := 0 ; i <= 10 ; i++ {
		go add(1,i)
	}
	
	time.Sleep(1e9)
	}


