package main

import (
	"strings"
	"fmt"
	"net/http"
)

func main()  {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request)  {
		fmt.Fprintf(w, "Hello World!!")
		r.ParseForm() //解析參數，預設是不會解析的
		//輸出到Server端的Message
		fmt.Println("path", r.URL.Path)
    	fmt.Println("scheme", r.URL.Scheme)
    	fmt.Println(r.Form["url_long"])
    	for k, v := range r.Form {
        	fmt.Println("key:", k)
        	fmt.Println("val:", strings.Join(v, ""))
    	}
	})

	http.ListenAndServe(":8080", nil)
}