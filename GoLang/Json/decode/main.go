package main

import (
	"fmt"
	"encoding/json"
)

type User struct {
	Name string
	ID string
	Age uint
	Merriage bool
	Skill []string
}

func main()  {
	usr := User{
		"MeatTaro",
		"M122229868",
		30,
		false,
		[]string{"coding", "hiking", "Music"},
	}
	u, err := json.Marshal(usr)
	if err != nil {
		fmt.Printf("false: %v\n", err)
		return
	}
	//-------------------------------
	var usr2 User
	err = json.Unmarshal(u, &usr2)
	if err != nil {
		fmt.Printf("false: %v\n", err)
		return
	} 
	fmt.Printf("Json:%#v\n", usr2)
}