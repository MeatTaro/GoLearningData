package main

import (
	// "runtime"
	"context"
	"fmt"
	"time"
	"sync/atomic"
)

func AddNum(a *int32, b int, deferFunc func())  {
	defer func() {
		deferFunc()
	}()
	
	for i:=0; ; i++{
		curNum := atomic.LoadInt32(a)
		newNum := curNum +1
		time.Sleep(time.Millisecond *200)
		// fmt.Printf("index:%d,goroutine Num:%d \n",i,runtime.NumGoroutine())
		if atomic.CompareAndSwapInt32(a, curNum, newNum) {
			fmt.Printf("number當前值: %d [%d-%d]\n", *a, b, i)
			break
		}else{
			fmt.Printf("The CAS operation failed. %d [%d-%d]\n",*a,b,i)
		}
		
	}
}

func main()  {
	total := 5
	var num int32
	fmt.Printf("number初始值: %d\n", num)
	fmt.Println("啟動子協程...")
	ctx, cancelFunc := context.WithCancel(context.Background())
	for i:=0; i<total; i++ {
		go AddNum(&num, i, func() {
			if atomic.LoadInt32(&num) == int32(total) {
				cancelFunc()
			}
		})
	}
	<- ctx.Done()
	fmt.Println("所有子協程執行完畢")
}