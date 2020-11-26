package main

import (
	"fmt"
	"sync"
)
//How to  use WaitGrout
	//1.A WaitGroup waits for a collection of goroutines to finish. 
	//2. The main goroutine calls Add to set the number of goroutines to wait for. 
	//3.Then each of the goroutines runs and calls Done when finished. 
	//4.At the same time, Wait can be used to block until all goroutines have finished.

var wg sync.WaitGroup //宣告type為WaitGroup的全域變數 (1. WaitGroup 用於等待一系列的goroutine完成任務)

var count int32

func AddOne() { //定義函數AddOne，每次調用時count加1
	defer wg.Done() //在程序結束前關閉goroutine(3.Done是一種goroutine計數器，負責關閉Add開啟的goroutine，須在其它新建的goroutine中呼叫)
	count++
}

func main()  {
	wg.Add(3) //往WaitGroup裡增設3個goroutine (2.Add是一種goroutine計數器，負責增設需要開啟的goroutine數量，需在main goroutine中呼叫)

	go AddOne()//調用3個goroutine
	go AddOne()
	go AddOne()

	wg.Wait()//阻塞 main goroutine直至所有 goroutine完成(4.Wait blocks until the WaitGroup counter is zero.)

	fmt.Printf("count: %d", count) //執行結束，輸出Count: 3
}



//參考資料: 
//https://ithelp.ithome.com.tw/articles/10226126
//https://golang.org/pkg/sync/#WaitGroup