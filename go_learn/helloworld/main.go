
//Go Package
//1. 所有程式皆由package構成
//2. 同一個目錄階層下的程式package 名稱需相同
package main

import (
	"helloworld/hello" //引入外部package
)

func main()  {
	// fmt.Println("hello world")
	hello.Hello() 
}

