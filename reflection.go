package main

import (
        "fmt"
         "reflect"
         )

// type User struct {
// 	Id int
// 	Name string
// 	Age int
// }


// func (u User) Hello() {
// 	fmt.Println("hello world")
// }


// func main() {
// 	u := User{1, "ok", 12}
// 	Info(u)

// }


// func Info(o interface{}) {
// 	t := reflect.TypeOf(o)
// 	fmt.Println("Type:", t.Name())

// 	v := reflect.ValueOf(o)
// 	fmt.Println("Fields:")

// 	for i := 0; i < t.NumField(); i++ {
// 		f := t.Field(i)
// 		val := v.Field(i).Interface()
// 		fmt.Printf("%s:%v = %v\n", f.Name,f.Type,val)
// 	}

// 	for i:= 0; i < t.NumMethod(); i++ {
// 		m := t.Method(i)
// 		fmt.Printf("%s:%v\n", m.Name,m.Type)
// 	}
// }


// func main() {
// 	x := 123
// 	v := reflect.ValueOf(&x)
// 	v.Elem().SetInt(999)/

// 	fmt.Println(v)
// }



type User struct {
	Id int
	Name string
	Age int
}

func main() {
	u := User{1 , "liu", 12}
	Set(&u)
	fmt.Println(u)
}

func Set(o interface{}) {
	v := reflect.ValueOf(o)

	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("xxx")
		return
	} else {
		v = v.Elem()

	}

	f := v.FieldByName("Name1")
	if !f.IsValid() {
		fmt.Println("ABC")
		return
	}

	if f.Kind() == reflect.String {
		f.SetString("BYBY")
	}
}