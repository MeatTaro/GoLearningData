package main

import (
	"net/http"
	"fmt"
)

func main() {
	//HandleFunc 設定訪問的路由"/"並註冊
	http.HandleFunc("/",func (w http.ResponseWriter, r *http.Request)  {
	//以匿名函數func(w http.ResponseWriter, r *http.Request{}建立接收/回傳方法)
		
		fmt.Fprintf(w, "Hello World") 	//輸出到客戶端的訊息
	})
	http.ListenAndServe(":8080", nil)//設置監聽的埠
	//ListenAndServe底層初始化一個server物件,呼叫（"tcp", addr），
	//也就是底層用tcp協議建立一個server，用來監聽埠
}

//參考資料：
//https://willh.gitbook.io/build-web-application-with-golang-zhtw/03.0/03.3