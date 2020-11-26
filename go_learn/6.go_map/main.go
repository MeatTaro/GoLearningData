package main

import (
	"fmt"
)

func main()  {
	m := make(map[string]string)

	m["go"] = "beego"
	m["python"] = "django"
	m["php"] = "laravel"
	m["ruby"] = "rails"

	delete(m,"ruby")
	
	for i, j := range m{
		fmt.Println(i ,":", j)
	}
	

	
}