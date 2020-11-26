package main

import (
	"sync"
	"time"
	"fmt"
)

func main()  {
	wg := new(sync.WaitGroup)
	n := 300000
	wg.Add(n) //加入需要啟動的goroutine數目
	start := time.Now().Unix()
	for i := 1; i <= n; i++ {
		go isPrime(i, wg) //傳入wg的值	
	}
	wg.Wait() //只要gorutine未歸零，就block住main函式
	end := time.Now().Unix()
	fmt.Println(end - start,"second")
}

func isPrime(n int, wg *sync.WaitGroup) {
	defer wg.Done() //defer 在函式結束前執行 wg.Done(減少gorouine)
	for i := 2; i < n; i++ {
		if n % i == 0 {
			return
		}	
	}
	fmt.Println(n)
}

