package main

import (
	"fmt"
	"sync"
)
//How to use Mutex
var counter int = 0

func add(a, b int, lock *sync.RWMutex) {
	c := a + b
	lock.Lock()
	counter ++
	fmt.Printf("%d : %d + %d = %d\n", counter, a, b, c)
	lock.Unlock()
}

func main()  {
	lock := &sync.RWMutex{}
	for i:=0; i<10; i++ {
		go add(1,i,lock)
	}

	for {
		// lock.RLock()
		c := counter
		// lock.RUnlock()
		if c >= 10 {
			break
		}

	}
	
}