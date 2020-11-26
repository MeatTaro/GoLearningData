package main

import (
	"fmt"
)

//standard: arrayName [arrayLength]arrayType {arrayName[0]= content0 arrayName[1]= content1}
//shorthand: arrayName := [arrayLength] arrayType {content0, content1,}
func main() {
	// var langs[4]string
	// 	langs[0]="go"
	// 	langs[1]="python"
	// 	langs[2]="php"
	// 	langs[3]="javascript"
	langs := [4]string{
		"go",
		"python",
		"php",
		"javascript",
	}
	for i,e := range langs{
		fmt.Printf("%d.%s\n",i+1, e)
	}

	var x string
	fmt.Println("whitch language do you want learn?")
	fmt.Scanln(&x)
	
	if x==langs[0]{
		fmt.Println( "hard" )
	}else if x==langs[1]{
		fmt.Println( "esey")
	}else if x==langs[2]{
		fmt.Println( "medum")
	}else if x==langs[3]{
		fmt.Println( "too easy")
	}else{
		fmt.Println( "languages doesn't exit")
	}
	
}



//學習重點
//1. array表示法
//1. array索引是從[0]開始，而非[1] 此陣列長度為4
//2. 練習條件語句用法
//3. 練習fmt.Scanln