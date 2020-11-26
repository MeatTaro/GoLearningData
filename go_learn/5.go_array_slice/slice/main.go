package main

import (
	"reflect"
	"fmt"
)

func main()  {
	langs := [4]string{
		"go",
		"python",
		"php",
		"javascript",
	}
	slice := langs[0:4]
	fmt.Println(reflect.TypeOf(langs))
	fmt.Println(reflect.TypeOf(slice))//slice本身不存值
	fmt.Println(langs[3])

	slice[3] = "C++" //array的值不可修改，但可透過修改slice來改變array的值
	fmt.Println(langs[3])

	matrix := [][]float64{
		[]float64{1,2,3,},
		[]float64{4,5,6,},
	}
	fmt.Println(reflect.TypeOf(matrix))
	fmt.Println(matrix[0][0],matrix[0][1],matrix[0][2])
	fmt.Println(matrix[1][0],matrix[1][1],matrix[1][2])
	
	movslice := make([]int,5)
	for i :=0; i<len(movslice); i++{
		n := i + 1
		movslice[i] = n*n
	}
	fmt.Println(movslice[0:5])

	movslice = append(movslice,7,8,9)
	for num, e := range movslice{
		fmt.Println(num,e)
	}

	movslice = append(movslice[0:4],movslice[5:6]...)
	fmt.Println(movslice)

	
}