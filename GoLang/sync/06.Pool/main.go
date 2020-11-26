package main

import (
	"fmt"
	"sync"
)

func testPut(pool *sync.Pool, deferfunc func()) {
	defer func ()  {
		deferfunc()
	}()
	value := "Hellow MeatTaro"
	pool.Put(value)
}

func main()  {
	var wg sync.WaitGroup
	wg.Add(1)
	var pool =&sync.Pool{
		New : func() interface{} {
			return "Hello,World"
		},
	}
	go testPut(pool, wg.Done)
	wg.Wait()
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
}