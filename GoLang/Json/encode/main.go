package main

import (
	"fmt"
	"encoding/json"
)

type User struct {
	Name string
	Website string
	Age  uint
	Male bool
	Skills []string
}

func main()  {
	usr := User{
		"MeatTaro",
		"https://google.com",
		30,
		true,
		[]string{"GoLang", "PHP", "Linux"},
	}
	u, err := json.Marshal(usr)
	if err != nil {
		fmt.Printf("fail : %v \n", err)
		return
	}
	fmt.Printf("Json: %s\n", u)
}