package main

import (
	"fmt"
)
type connection struct{
	email string
	phone string
}
type personal struct{
	name string
	age string
}

func main() {
	
	cnt := make(chan connection,1)
	connect(cnt)

	per := make(chan interface{},2)
	per <- personal{
		name: "MeatTaro",
		age: "31",
	}
	per <- "Hello MeatTaro"

	
	fmt.Println(<-per)
	fmt.Println(<-per)
	fmt.Println(<-cnt)
	// fmt.Println(cnt,per)


}

func connect(cnt chan connection)  {
	c := connection{
		email: "quasimodo9868@gmail.com",
		phone: "0988859858",
	}
	cnt <- c
}