package main

import (
	"fmt"
	"strings"
)

func main()  {
	a := "gopher"
	b := "hello world"
	//Compare 函數用於比較兩字符串大小
	fmt.Println(strings.Compare(a,b)) //a<b output -1
	fmt.Println(strings.Compare(a,a)) //a=a output 0
	fmt.Println(strings.Compare(b,a)) //b>a output 1
	//Join用於連接2字串
	var s []string
	s = append(s, a, b)
	fmt.Println(strings.Join(s, "&"))

}