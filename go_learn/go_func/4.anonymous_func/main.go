package main

import "fmt"

//anonymous func(常使用在go routine)

func main() {
	//變數宣告 - 賦值為函式 - 定義函式類型 - 回傳值 - 輸出函式
	foo := func() string { 
		return "Hello World1"
	}

	fmt.Println(foo())
//---------------------------------------------------------------
	bar := func() {					//省略定義類型
		fmt.Println("Hello World2")	//回傳值改為直接輸出內容
	}

	bar()							//輸出函式
//---------------------------------------------------------------
	func() {						//省略變數宣告 
		fmt.Println("Hello World3") 
	}()								//直接執行func

//------------------------------------------------------go routine
	go func(i, j int) { 
		fmt.Println(i + j)
	}(1, 2)
//go func (go routine) 在程式尚未完成就輸出，所以不產生執行結果
}