package main

import (
	"sync"
)
//如何宣告WaitGroup
//A WaitGroup must not be copied after first use.
func main()  {
	var wg1 sync.WaitGroup
	wg2 := new(sync.WaitGroup)
	wg3 := &sync.WaitGroup{}
	var wg4 = &sync.WaitGroup{}
}