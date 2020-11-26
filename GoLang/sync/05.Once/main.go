package main

import (
	"time"
	"sync"
	"fmt"
)

func dosomething(o *sync.Once)  {
	fmt.Println("start:")
	o.Do(func ()  {
		fmt.Println("do something")
	})
	fmt.Println("end")
}

func main()  {
	o := &sync.Once{}
	go dosomething(o)
	go dosomething(o)
	time.Sleep(time.Second*1)
}