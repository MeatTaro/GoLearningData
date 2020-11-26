package main

import (
	"fmt"
)

func main()  {
	var h Human
	//此處可看出 Being嵌入Human ， Human嵌入Student 嵌入後卻又保持
	s := Student{Grade:1, Major:"English", Human: Human{Name:"MeatTro", Age:31, Being: Being{IsLive:true}}}

	fmt.Println("student:", s)
	fmt.Println("student:", s.Name, ", isLive:", s.IsLive, ", age:", s.Age, ", grade:", s.Grade, ", major:", s.Major)
	
	fmt.Println(h)
	fmt.Println(s.Human.Being)

	Heal(s.Human.Being)
	
	//透過s.呼叫了Student的Drink方法，但因為Student並未implement Eat()，所以這邊的s.Eat()是透過Human呼叫。
	//由於h.Eat()裡利用h.Drink()呼叫Drink()，所以結果會產生"human drinking ..."
	s.Drink()
	s.Eat()
}

type Being struct {
	IsLive bool
}

type Human struct {
	Being
	Name string
	Age int
}

func (h Human) Drink() {
	fmt.Println("human drinking ...")
}

func (h Human) Eat() {
	fmt.Println("human eating ...")
	h.Drink()
}

func (h Human) Move() {
	fmt.Println("human moving ...")
}

type Student struct {
	Human
	Grade int
	Major string
}

func (s Student) Drink() {
	fmt.Println("student drinking ...")
}

// type IEat interface {
//     Eat()
// }
 
// type IMove interface {
//     Move()
// }
 
// type IDrink interface {
//     Drink()
// }

func Heal(b Being){
    fmt.Println(b.IsLive)
}