package main

import (
	"fmt"
)

type point struct{
	x,y int
}

func main()  {
	
	p := point{1, 2} 

//General

	//輸出 初始化格式的 值
	fmt.Printf("%v\n", p) // {1,2}

	//輸出 初始化格式包含成員、值
	fmt.Printf("%+v\n", p) //{x:1 y:2}

	//輸出 初始化格式的值的 Go語法表示方式
	fmt.Printf("%#v\n", p) //main.point{x:1,y:2}

	//輸出 初始化格式的值的類型
	fmt.Printf("%T\n", p) //main.point
//Boolean
	//輸出 初始化格式的值的布林變量
	fmt.Printf("%t\n", true) //true

//Integer
	//輸出 整數的10進位表示方式
	fmt.Printf("%d\n", 123) //123

	//輸出 整數的2進位表示方式
	fmt.Printf("%b\n", 123) //1111011

	//輸出 整數的16進位表示方式
	fmt.Printf("%x\n", 123) // 7b

	//輸出 整數的字符表示方式
	fmt.Printf("%c\n", 123) // {

//float
	//輸出 浮點數初始格式化
	fmt.Printf("%f\n", 78.9) //78.900000

	// `%e`和`%E`使用科學計算方法來輸出整數
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)

//string
	//輸出 初始化格式的字符字串
	fmt.Printf("%s\n", " \"string\" ")

	//輸出 go源碼的字符字串
	fmt.Printf("%q\n", " \"string\" ")
}