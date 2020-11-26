package main

import (
	"time"
	"fmt"
	"sync"
)

var sharedrsc = false

func main()  {
	var wg sync.WaitGroup
	wg.Add(2)

	//Cond implements a condition variable, a rendezvous point for goroutines waiting for or announcing the occurrence of an event.
	//Each Cond has an associated Locker L (often a *Mutex or *RWMutex), which must be held when changing the condition and when calling the Wait method.


	//Cond實現了一種條件變數，供多個goroutine作為匯合點來等待通知事件發生。
	//每一個Cound必須關聯一個*Mutex or *RWMutex，當修改condition or 呼叫method Wait 時來加鎖保護condition

	m := sync.Mutex{}//初始化一個Mutex互斥鎖
	c := sync.NewCond(&m)//NewCond returns a new Cond with Locker l.
	//創建一個新的條件變數，並關聯互斥鎖m

	// this go routine wait for changes to the sharedrsc
    // 在sharedrsc被修改前，持續等待
	go func() {
		c.L.Lock()
		if sharedrsc == false {
			fmt.Println("goroutine 1 wait")
			c.Wait()
			//c.Wait暫停被呼叫的goroutine並自動解鎖c.L.Lock。(解鎖期間可供外部修改全局變量)接著在恢復執行goroutine前，再度鎖上
		}
		fmt.Println("gproitine 1", sharedrsc)
		c.L.Unlock()
		wg.Done()
	}()

	go func ()  {
		c.L.Lock()
		if sharedrsc == false {
			fmt.Println("goroutine 2 wait")
			c.Wait()
		}
		fmt.Println("goroutine 2 ",sharedrsc)
		c.L.Unlock()
		wg.Done()
	}()

	time.Sleep(2*time.Second)
	c.L.Lock()
	fmt.Println("main goroutine ready")
	sharedrsc = true
	c.Broadcast() //喚醒所有等待中的goroutine
	fmt.Println("main goroutine broadcast")
	c.L.Unlock()
	wg.Wait()//4. Wait 方法可以阻塞等待main goroutine直至所有 goroutine完成(blocks until the WaitGroup counter is zero.)
}

//參考資料: 
//https://www.jishuwen.com/d/2Xd2/zh-tw
//https://golang.org/pkg/sync/#Cond