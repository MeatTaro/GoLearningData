package main

import (
	"os"
	"strconv"
	"fmt"
	"calc/simplemath"
)
//定義一個用於印出程序使用說明的函數
var Usage = func() {
	fmt.Println("USAGE: calc command [arquments] ...")
	fmt.Println("\nThe commands are:\n\tadd\t計算兩個數值相加\n\tsqrt\t計算一個非負數平方根")
}

func main()  {

//os.Args用於獲得TERINAL寫入的參數，注意程序名本身是第一个参数
//比如 go run calc.go add 1 2 这条指令，第一个参数是 main

	args := os.Args
//除程序名本身外，至少需要传入两个其它参数，否则退出
	if args == nil || len(args) < 3 {
		Usage()
		return
	}

	// 第二个参数表示计算方法
	switch args[2] {
		// 如果是加法的话
		case "+":
			// 至少需要包含四个参数
			if len(args) != 4 {
				fmt.Println("USAGE: calc add <integer1><integer2>")
				return
			}
			// 获取待相加的数值，并将类型转化为整型
			v1, err1 := strconv.Atoi(args[1])
			v2, err2 := strconv.Atoi(args[3])
			//若參數讀取錯誤，則回報錯誤並退出
			if err1 != nil || err2 != nil {
				fmt.Println("USAGE: ccalc add <integer1><integer2>")
				return
			}
			//從自定義package simplemath 導入Add func 進行加法計算
			ret := simplemath.Add(v1,v2)
			fmt.Println("Result", ret)
		case "-^":
			if len(args) != 3 {
				fmt.Println("USAGE: ccalc -^ <integer1>")
				return
			}
			v, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("USAGE: ccalc -^ <integer1>")
				return
			}
			ret := simplemath.Sqrt(v)
			fmt.Println("Result", ret)
		default:
			Usage()		
	}
}

