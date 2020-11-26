package main

import (
	"time"
	"fmt"
)

func add(x, y int, ch chan int) {
	z := x + y
	fmt.Printf("%d + %d = %d\n",x,y,z)
	ch <- 1
}

func main() {
	start := time.Now()
	chs := make([]chan int, 10)
	for i:=0; i<10; i++ {
		chs[i] = make(chan int)
		go add(1,i,chs[i])
	}
	for _, ch := range chs {
		<- ch
	}
	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println(consume)
}