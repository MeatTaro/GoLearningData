package main

import (
	"fmt"
)

func makeEvenGenerator()func() int { //makeEvenGenerator() 回傳值為一函式func() type為int
	i:= 0 		//宣告一區域變數
	return func()(ret int){			//此處return給makeEvenGenerator()的值為func()，並定義func回傳ret type為int。
		ret = i
		i+=2
		return  //return ret的結果
	}

}
func main()  {
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
}