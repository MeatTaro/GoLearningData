package main

import (
	"fmt"
)

func test(c, q chan int) {
	i := 0
	for {
		select {
			case c <- i:  //i值 send to channel c成立，就執行i=i+1
				i = i + 1
			case <-q :
				fmt.Println("quit") //若channel q 存在值(0)，receive並印出
				return
		}
	}
}

func main()  {
	c := make(chan int)
	q := make(chan int)
	//當channel c 存在值(from test())，receive值並印出，直到迴圈i>5 ，send 0 進入channel q
	go func(){
		for i:=0; i<5; i++ {
			fmt.Println(<- c) 
		}
		q <- 0
	}()
	test(c,q)
}