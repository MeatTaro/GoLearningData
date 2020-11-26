package main

import (
	"strconv"
	"io"
	"time"
	
	
	"crypto/md5"
	
	"log"
	"html/template"
	"fmt"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Method:", r.Method)
	
	switch r.Method {
	case "GET":
		crutime := time.Now().Unix()
		//md5.New()返回一個hash interface並初始化h 
		h := md5.New()
		//strconv.FormatInt將時間戳格式化為10進位整數字符串
		//io.WriteString以hash interface作為緩衝區將字符串轉化為hash
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		//h.Sum返回hash interface 底下的資料
		token := fmt.Sprintf("%x", h.Sum(nil))
		fmt.Println(token)
		t, _ := template.ParseFiles("login.html")
		log.Println(t.Execute(w, token))
	case "POST":	
		r.ParseForm()
		m := r.Form.Get("token")
		if m != "" {
			fmt.Println(m)
		} else {
			fmt.Println("error")
		}
		
        fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) //輸出到伺服器端
		template.HTMLEscape(w, []byte(r.Form.Get("username"))) //輸出到客戶端
		if len(r.Form["username"][0]) == 0{
			fmt.Fprintf(w, "username can not nil")
			fmt.Println("username can not nil")
		}

		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
        
	}
}
func main()  {
	http.HandleFunc("/login", login)
	http.ListenAndServe(":8080", nil)
}