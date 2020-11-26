package main

import (
	// "time"
	"fmt"
)

func main()  {
	//1.channel are a typed conduit through which you can send and receive values with the channel operator <-
	var chanInt chan int //宣告var名稱，channel 傳遞類型

	//2.like maps and slices, channels must be created before use.
	
	chanInt = make(chan int,3)//使用make為channel建立容器 ,容量為3
	//channel can be buffered. Provide the buffer length as the argument to 'make' to  initialize a buffered channel

	//3.By default, sends and receives block until the other side is ready.
	chanInt <- 1 //send 值 進 channel
	chanInt <- 2
	chanInt <- 3
	//Sends to a buffered channel block only when the buffer is full.
	fmt.Printf("channel容量:%v element長度:%v\n",cap(chanInt),len(chanInt)) 
	//check! 印出channel的容量,長度
	fmt.Println(<- chanInt) 
	//receive value in channel 裡的值 並且採取先進先出
	fmt.Printf("channel容量:%v element長度:%v\n",cap(chanInt),len(chanInt))//receive vale 前後容量與長度的變化

	//4.若未給定容量(nil)，必須於gorutine執行
	go func() {
		chanString := make(chan string)
		chanString <- "hello world"
		fmt.Println(<- chanString)
	}()

	ch := make(chan int, 5) //initialize a bufferd channel  
	go fibonacci(cap(ch), ch)
	fmt.Printf("channel容量:%v element長度:%v\n",cap(ch),len(ch))
	for i := range ch {
		//此處透過range接通channel，依序receive每個值
		fmt.Println(i)
	}

	_, ok := <-ch 
	fmt.Println(ok)
}

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i:=0; i<n; i++ {
		ch <- x
		x, y = y, x+y
		//fmt.Printf("channel容量:%v element長度:%v\n",cap(ch),len(ch))
		//check! value is sending to channel
		//channel容量:5 element長度:0
		// channel容量:5 element長度:1
		// channel容量:5 element長度:2
		// channel容量:5 element長度:3
		// channel容量:5 element長度:4
	}
	
	close(ch)
	//由於主協程不知道子協程(fibonacci())何時執行完畢，若持續從一個空的channel receive values，會導致deadlock
	//close()不可重複執行，否則會引發panic
	//close()必須由 sender執行，因為receiver不會知道什麼時候發送執行完畢
}




