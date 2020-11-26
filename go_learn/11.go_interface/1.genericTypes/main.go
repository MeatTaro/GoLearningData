package main

import (
	"fmt"
)

func Switch(input interface{}) {

	switch v := input.(type) {
	case float64:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*3)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}

}

func main() {
    //這裡interface作為泛用型別
	var a interface{}
	var b interface{}
	var c interface{}

	a = 5
	b = 1.5
    c = "hello world"
    //經過func Switch 分別印出a,b,c經定義後的型別
	Switch(a)
	Switch(b)
	Switch(c)
	fmt.Println(a, b, c)
}
