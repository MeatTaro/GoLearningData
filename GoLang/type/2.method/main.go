package main

import (
	"fmt"
)

type Person struct { 
	id uint
	name string
	male bool
	age int
}

func NewPerson(id uint, name string, male bool, age int) *Person {
	return &Person{id, name, male, age}
}

//method為一種特殊函數，用於define receiver(type的變數)
//how to add method to type
//func (type宣告) method()命名
func (x Person) GetName() string {
	return x.name
}

//if the value of field belong type have to be set, the way we call the type in method is point. but the parameter's type in method doesn't need to be setted, so it's not point. 
func (x *Person) SetName(name string)  {
	x.name = name
}

func main()  {
	//after initialize type 'Person', useing method 'GetName()' to get the value of 'name' and print out
	person := NewPerson(1, "MeatTaro", true, 31)
	fmt.Println(person.GetName())
	//using method 'SetName' to set the value of 'name' and print out
	person.SetName("Christian")
	fmt.Println(person.GetName())
}