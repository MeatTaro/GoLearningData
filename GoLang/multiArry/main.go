package main

import (
	"fmt"
)

func main()  {
	var multi [9][9]string
	for j := 0; j < 9; j++ {
		for i := 0; i < 9; i++ {
				n1 := i + 1
				n2 := j + 1

				multi[i][j] = fmt.Sprintf("%dx%d=%d", n2, n1, n1 * n2)
				// fmt.Println(multi[i][j]) //依2迴圈規則印出
			}
	  }
	  for _, v1 := range multi {
		// fmt.Println(v1) //依第一層迴圈印出所有1維陣列與陣列中的二維值
		for _, v2 := range v1 {
			fmt.Println(v2) //依二維陣列中的值依序印出
			fmt.Printf("%-8s", v2)  // 位宽为8，左对齐
		}
		fmt.Println()
	}
	slice0 := make([]int, 5, 10) 
	fmt.Println(slice0)
}