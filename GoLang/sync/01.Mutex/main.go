package main

import (
	"fmt"
	"sync"
)
//How to use Mutex
//一但數據被多個routine共享，那就有機會產生競用和衝突(race condition)。為了避免多個routine同時操作or修改一個數據模塊(共享資源)，可以使用互斥鎖(Mutex)協調數據模塊的訪問權限，確保同一時間僅有一個routine能夠操作or修改數據模塊。

var counter int = 0 //宣告一個全域變數(數據模塊)

func add(a, b int, lock *sync.Mutex) { //定義函式傳入一個type為*sync.Mutex 的接口 
	c := a + b
	lock.Lock() //透過接口lock引入package內的Lock and Unlock函數
	counter ++
	fmt.Printf("%d : %d + %d = %d\n", counter, a, b, c)
	lock.Unlock() //互斥鎖區間就像一個routine流量控制閘，每次僅允許一個goroutine於區域內進行操作。

	//不要重复锁定互斥锁；
	//不要忘记解锁互斥锁，必要时使用 defer 语句；
	//不要对尚未锁定或者已解锁的互斥锁解锁；
	//不要在多个函数之间直接传递互斥锁。
}

func main()  {
	lock := &sync.Mutex{} // main goroutine中宣告並初始化一個變數，作為互斥鎖

	for i:=0; i<10; i++ { //利用條件迴圈開啟10個goroutine
		go add(1,i,lock)
	}
	
	//雖然我們利用Mutex限制了goroutine的流量，但goroutine執行速度仍快於顯示速度

	//利用for迴圈提供條件循環，直到counter計數器 >= 10才結束程序
	for {
		lock.Lock()
		c := counter
		lock.Unlock()
		if c >= 10 {
			break
		}

	}
	
}