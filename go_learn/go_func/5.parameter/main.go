package main

import (
	"strings"
	"fmt"
)

// func getUserListSQL(username, email string) string  {
// 	sql := "selsct from user"
// 	where := []string{}

// 	if username != ""{
// 		where = append(where, fmt.Sprintf("username = '%s",username))
// 	}
// 	if email != ""{
// 		where = append(where, fmt.Sprintf("email = '%s",email))
// 	}
// 	return sql + "where" + strings.Join(where, " or ")
// }

type searchOpt struct{
	username string
	email string		
}

func getUserListoptsSQL(opts searchOpt) string  {
	sql := "selsct from user"
	where := []string{}

	if opts.username != ""{
		where = append(where, fmt.Sprintf("username = '%s",opts.username))
	}
	if opts.email != ""{
		where = append(where, fmt.Sprintf("email = '%s",opts.email))
	}
	return sql + "where" + strings.Join(where, " or ")
}
func main()  {
	// fmt.Println(getUserListSQL("MeatTaro", ""))
	// fmt.Println(getUserListSQL("MeatTaro", "test@gmail,.com"))

	fmt.Println(getUserListoptsSQL(searchOpt{
		username:	"MeatTaro",
		email: 		"test@gmail",
	}))
}