package main

import (
	"time"
	"fmt"
)

func main()  {
	start := time.Now()
	n := 100000
	for i := 1; i <= n; i++ {
		go isPrime(i)		
	}
	var x int
	fmt.Scanln(&x)
	end := time.Now()
	fmt.Println(end.Unix() - start.Unix(),"second")
}

func isPrime(n int) {
	for i := 2; i < n; i++ {
		if n % i == 0 {
			return
		}	
	}
	fmt.Println(n)
}