package main

import (
	    "html/template"
	    "net/http"
	)

type Person struct {
    Name string // 命名要大寫開頭
}
//Handler
func indexHandle(w http.ResponseWriter, r *http.Request) {
    t,_ := template.ParseFiles("index.html") // 讀取 HTML 檔案
    p := Person{Name:"MeatTaro"} // 要送到前端的訊息
    t.Execute(w, p)
}

func main()  {
	http.HandleFunc("/", indexHandle) //當瀏覽器輸入根目錄
	http.ListenAndServe(":8080",nil)
}