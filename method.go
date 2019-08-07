package main

import (
	"fmt"
)

// type A struct {
// 	Name string
// }

// type B struct {
// 	Name string
// }


// func main() {
// 	a := A{}
// 	a.Print()
// 	fmt.Println(a.Name)
	
// 	b := B{}
// 	b.Print()
// 	fmt.Println(b.Name)
// }

// func (a *A) Print() {
// 	a.Name = "AA"
// 	fmt.Println("A")
// }

// func (b B) Print() {
// 	b.Name = "BB"
// 	fmt.Println("B")
// }



// type TZ int

// func main() {
// 	var a TZ
// 	a.Print()    //method value
// 	(*TZ).Print(&a)  //method expression
// }

// func (a *TZ) Print() {
// 	fmt.Println("TZ")
// }

// type A  struct {
// 	name string  //小写与大写字段的区别是以包为依据，对于同一个包都是公有字段，而小写对于不同包则为私有字段
// }

// func main() {
// 	a := A{}
// 	a.Print()
// }

// func (a *A) Print() {
// 	a.name = "123"
// 	fmt.Println(a.name)
// }

type A int

func main() {
	var a A
	a.Increase(100)
	fmt.Println(a)

}

func (a *A) Increase(num int) {
	*a += A(num)
	fmt.Println(*a)
}