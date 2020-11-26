package main

import (
	"fmt"
	"runtime"
)

func sum(s int, ch chan int)  {
	sum := 0
	for i:=1; i<=100; i++{
		sum += i
	}
	fmt.Printf("子協程%d運算結果:%d\n",s,sum)
	ch <- sum
}

func main()  {
	cpus := runtime.NumCPU()
	chs := make([]chan int,cpus)
	for i:=0; i<len(chs); i++ {
		chs[i] = make(chan int, 1)
		go sum(i,chs[i])
	}
	sum := 0
	for _, ch := range chs{
		res := <- ch
        sum += res
	}
	fmt.Printf("最終運算結果:%d",sum)
}