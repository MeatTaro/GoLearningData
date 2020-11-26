package main

import (
	"fmt"
)
//請將以下程式以物件導向方式重新撰寫，並利用Interface建立相關物件
// type AnimalsType int

// 	const (
// 		Duck AnimalsType = iota
// 		Dog
// 		Tiger
// 	)

// func New(t AnimalsType) string {
// 	switch t {
// 	case Duck:
// 		return "Pack Pack"
// 	case Dog:
// 		return "Wow Wow"
// 	case Tiger:
// 		return "Halum Halum"
// 	default:
// 		panic("Unknow animals type")
// 	}	
// }
// func main()  {
// 	animals := make([]AnimalsType, 0)
// 	animals = append(animals,Duck ,Dog ,Tiger)
// 	for _, a := range animals{
// 		fmt.Println(New(a))
// 	}
// }

type AnimalsType int
const (
	Duck AnimalsType = iota
	Dog
	Tiger
)
type DuckClass struct{}

func NewDuck() *DuckClass  {
	return new(DuckClass)
}
func (t *DuckClass) Speak() {
	fmt.Println("Pack Pack")
}
type DogClass struct{}

func NewDog() *DogClass {
	return new(DogClass)
}
func (t *DogClass) Speak() {
	fmt.Println("Wow Wow")
}

type TigerClass struct{}

func NewTiger() *TigerClass{
	return new(TigerClass)
}
func (t *TigerClass) Speak() {
	fmt.Println("Halum Halum")
}

type IAnimals interface{
	Speak()
}
func New(t AnimalsType) IAnimals {
	switch t {
	case Duck:
		return NewDuck()
	case Dog:
		return NewDog()
	case Tiger:
		return NewTiger()
	default:
		panic("Unknow animals type")
	}
}
func main()  {
	// animals := make([]IAnimals, 0)
	duck := New(Duck)
	dog := New(Dog)
	tiger := New(Tiger)
	fmt.Printf("%T,%T,%T\n", duck,dog,tiger)
	//duck,dog,tiger為初始化後的Class
	// animals = append(animals,duck ,dog ,tiger)
	animals := []IAnimals{duck,dog,tiger}
	fmt.Printf("%T\n", animals)
	for _, a := range animals{
		a.Speak()
	}
}