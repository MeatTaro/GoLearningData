package main

import (
	"time"
	"fmt"
)

func main() {
    n := 30000 //定義目標整數範圍
    start := time.Now() //紀錄起始時間
    for i:=1;i<=n;i++{ //利用for迴圈 在目標整數n範圍內傳入所有整數i進入isPrime函數(質數判斷)
        if isPrime(i){ //如果isPrime(i)有回傳值，則印出i
            fmt.Println(i)
        }
    }
    end := time.Now() //紀錄結束時間
    fmt.Println(end.Unix() - start.Unix(), "second") //印出時間差
}
func isPrime(n int) bool { //建立函數，傳入type為int變數n、回傳值type為bool
	for i:=2;i<n;i++{ //成立一迴圈，變數i賦值為2(初始化)，條件為i<n，符合條件則執行i++
		if n % i == 0 {
			return false
		}
	}
	return true
}