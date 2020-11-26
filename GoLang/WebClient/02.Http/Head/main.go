package main

import (
	"fmt"
	"net/http"
)

func main()  {
	resp, err := http.Head("https://xueyuanjun.com")
	if err != nil {
		fmt.Println("fail", err.Error())
		return
	}

	defer resp.Body.Close()

	for key, value := range resp.Header {
		fmt.Println(key, ":", value)
	}
}