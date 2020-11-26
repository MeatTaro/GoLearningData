package main

import (
	"log"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		params := request.URL.Query()
		fmt.Fprintf(writer, "Hello! %s", params.Get("name"))
	})
	err := http.ListenAndServeTLS(":8080", "./ca.crt", "./ca.key",nil)
	if err != nil {
		log.Fatal("Fail: %v", err)
	}
}