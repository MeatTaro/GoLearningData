package main

import (
	"fmt"
	"errors"
)

func add(a, b int) (c int, err error)  {
	if a<0 || b<0 {
		err = errors.New("XXXXX")
		return
	}
	a *= 2
	b *= 3
	c = a + b
	return
}

func main()  {
	a, b := 1, 2
	c, err := add(a,b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("add(%d,%d)=%d\n",a,b,c)
	}
}