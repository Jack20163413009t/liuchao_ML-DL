package main

import "fmt"

// func main() {
// 	var fs = [4]func() {}
// 	for i:=0;i<4;i++ {
// 		defer fmt.Println("defer i = ",i)
// 		defer func() {fmt.Println("defer_closure i=",i)} ()
// 		fs[i] = func() {fmt.Println("closure i=",i)}

// 	}
// 	for _,f := range fs {
// 		f()
// 	}
// }

// type person struct{
// 	Name string
// 	Age int
// }

// func main() {
// 	a := &person{
// 		Name:"liuchao",
// 		Age:18,
// 	}

// 	// a.Name = "liuchao"
// 	// a.Age = 18
// 	a.Name = "du"
// 	fmt.Println(a)
// 	A(a)
// 	B(a)
// 	fmt.Println(a)
// }

// func A(per *person) {
// 	per.Age = 16
// 	fmt.Println("A",per)
// }

// func B(per *person) {
// 	per.Age = 15
// 	fmt.Println("B",per)
// }

// func main() {
// 	a := &struct {
// 		Name string
// 		Age int
// 	} {
// 		Name:"liu",
// 		Age:18,
// 	}
// 	fmt.Println(a)
// }



// type person struct {
// 	Name string
// 	Age int
// 	Contact struct {
// 		Tel,Adr string
// 	}
// }

// func main() {
// 	a := person{ Name:"liu",
// 	Age:16,
// }
// 	a.Contact.Tel = "911"
// 	a.Contact.Adr = "chengdu"

// 	fmt.Println(a)
// }



type human struct {
	Sex int
	Age int
}

type teacher struct{
	human
	Name string
	Age int
}

func main() {
	a := teacher{human:human{Sex:0,Age:10}, Name:"liu", Age:16,}
	// a.Sex = 10
	a.human.Age = 15
	fmt.Println(a)
}