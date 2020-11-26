package main

//HTTP構成的網路應用為
//1. Client端發出request
//2.Server端接受request
//3.Server回傳response
//4. Client端接收並處理response

//當Server端接收到request時先進入router，router會找到對應的處理器(Handler)1u/4對request進行處理並構建response

import (
	"fmt"
	"net/http"
)

//type HandlerFunc func(ResponseWriter, *Request)

// func main()  {
// 	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request)  {
// 		fmt.Fprintf(w, "Hello World")
// 	})
// 	http.ListenAndServe(":8080", nil)
// }

//func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)

type indexHandler struct {
	content string
}

func (i *indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, i.content)
}

func main()  {
	http.Handle("/", &indexHandler{content: "Hello World XD"})
	http.ListenAndServe(":8080", nil)
}

//http.HandleFunc和http.Handle都是用于注册router
//以上方法最後都呼叫ListenAndServe開啟監聽
//func ListenAndServe(addr string, handler Handler) error
//addr表示監聽的IP地址端口號, handler表示Server端的處理程序，通常為nil，表示調用default http.DefaultServeMux


//參考資料:
//https://juejin.im/post/6844903998869209095
//https://golang.google.cn/pkg/net/http/#Handler