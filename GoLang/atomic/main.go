package main

import (
	"sync/atomic"
	"fmt"
)

func main() {
	//AddInt32 atomically adds delta to *addr and returns the new value.
	var i int32 = 1
	atomic.AddInt32(&i,1)
	fmt.Println("i=i+1=",i)
	atomic.AddInt32(&i,-1)
	fmt.Println("i=i-1=",i)
	//CompareAndSwapInt32 executes the compare-and-swap operation for an int32 value.
	//func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool)
	var a int32 = 1
	var b int32 = 2
	var c int32 = 2
	d := atomic.CompareAndSwapInt32(&a, a, b)
	fmt.Println("a, b, c, d:", a, b, c, d)
	e := atomic.CompareAndSwapInt32(&b, b, c)
	fmt.Println("a, b, c, d, e:", a, b, c, d, e)
	//LoadInt32 atomically loads *addr.
	//StoreInt32 atomically stores val into *addr.
	var x int32 = 100
	var y int32
	atomic.StoreInt32(&y, atomic.LoadInt32(&x))
	fmt.Println("x,y:",x,y)

	var j int32 = 0
	var k int32 = 1
	fmt.Printf("(%d , %d)\n",j,k)
	l := atomic.SwapInt32(&j, k)
	fmt.Printf("(%d , %d)\n",j,k)
	fmt.Println(l)

	var z atomic.Value
	z.Store(100)
	fmt.Println(z.Load())
}