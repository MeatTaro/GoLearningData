package main

import (
	"log"
	"html/template"
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello World")
}

func login(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		//template.ParseFiles("")創建模板回傳到t 並解析文本""
		t, _ := template.ParseFiles("login.html")
		//t.Execute執行解析後的模板，執行過成為合併、替換文本的{{.}}
		log.Println(t.Execute(w, "HelloWorld"))
	} else {
		r.ParseForm()//server端解析資料表單
		fmt.Println("username:", r.Form["username"])
		//r.Form 接收來自客戶端,html定義,符合 "username" "password"的資訊
		fmt.Println("password:", r.Form["password"])
	}
}

func main()  {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}