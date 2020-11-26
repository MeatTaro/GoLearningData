package main

import (
	"regexp"
	"strconv"
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
	switch  r.Method {
	case "GET":
		t, _ := template.ParseFiles("login.html")
		log.Println(t.Execute(w, nil))
	case "POST":
		r.ParseForm()

		//使用者名稱
		// fmt.Println("username:", r.Form["username"])
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username")))
		if len(r.Form["username"][0]) == 0{
			fmt.Fprintf(w, "username can not nil")
			fmt.Println("username can not nil")
		}
		
		//密碼
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))

		//真實姓名
		fmt.Println("realname", r.Form["realname"])
		m, _ := regexp.MatchString("^\\p{Han}+$", r.Form.Get("realname"))
		if m != true {
			fmt.Fprintf(w, "please input chinese")
			fmt.Println("please input chinese")
		}
		//年齡
		fmt.Println("age:", r.Form["age"])
		getint, err := strconv.Atoi(r.Form.Get("age"))
		if err != nil {
			fmt.Fprintf(w, "age should be int")
			fmt.Println("age should be int")
		} else if getint > 100 {
			fmt.Fprintf(w, "too big")
			fmt.Println("too big")
		}
		//下拉選單
		slice := []string{"apple","pear", "banana"}
		v := r.Form.Get("fruit")
		for _, item := range slice {
			if item == v {
				// fmt.Fprintf(w, "not fruit")
				// fmt.Println("not fruit")
				fmt.Println("fruit", r.Form["fruit"])
			}
			
		}
		
	}
	
}

func main()  {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/login", login)
	http.ListenAndServe(":8080", nil) 
}