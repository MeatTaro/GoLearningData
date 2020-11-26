package main

import (
	"fmt"
)

//how to define "type"
//Person is the name of this type, define with a struct.
//the struct have 4 fields
type Person struct { 
	id uint
	name string
	male bool
	age int
}
//how to initialize a type
//使用大寫字母字首命名函數名稱，make a initialization function
//take fields be parameters to initialize type named Person, and return point of the type
// func NewPerson(id uint, name string, male bool, age int) *Person {
// 	return &Person{id, name, male, age}
// }
//初始化特定字串
func NewPerson(name string) *Person {
	return &Person{name:name}
}
//未經過初始化函數的變數皆會自動初始化為該類型的零值
func main()  {
	person := NewPerson("MeatTaro")
	fmt.Println(person)
}
//func main 驗證是否初始化成功