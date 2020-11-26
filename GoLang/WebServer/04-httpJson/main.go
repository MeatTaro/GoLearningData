package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type UserInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	http.HandleFunc("/api/query", func(w http.ResponseWriter, r *http.Request) {
		u := &UserInfo{
			Name: "syhlion",
			Age:  18,
		}
		b, err := json.Marshal(u)
		if err != nil {
			log.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}