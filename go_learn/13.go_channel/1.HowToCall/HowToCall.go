package main

import (
	"fmt"
)

func main()  {
	var chanInt chan int //宣告變數為channel，其中可傳入的值型態須為int
	chanInt = make(chan int,3) //為channel建立容器 ,容量為1
	chanInt <- 1 //send 值 進 channel
	chanInt <- 2
	chanInt <- 3
	fmt.Printf("channel容量:%v element長度:%v\n",cap(chanInt),len(chanInt)) //check! 印出channel的容量,長度
	fmt.Printf("channel容量:%v element長度:%v\n",cap(chanInt),len(chanInt))
	fmt.Println(<- chanInt) //取出 channel 裡的值 採取先進先出
	fmt.Printf("channel容量:%v element長度:%v\n",cap(chanInt),len(chanInt))
}