package main

import (
"fmt"
// "sort"
)

// func main() {

    // var m map[int]map[int]string
    // m[1] = make(map[int]string)
    // m[1][1] = "ok"
    // a := m[1][1]
    // fmt.Println(a)

        // var m map[int]map[int]string
        // m = make(map[int]map[int]string)
        // a,ok := m[1][1]
        // if !ok {
        //     m[1] = make(map[int]string)

        // }
        // m[1][1] = "good"
        // a,ok = m[1][1]
        // fmt.Println(a,ok)

    // for i,v:=range slice{
    //     slice[i]
    // }

    // sm := make([]map[int]string,5)
    // for _,v := range sm {
    //     v = make(map[int]string,1)
    //     v[1] = "ok"
    //     fmt.Println(v)
    // }
    // fmt.Println(sm)

    // sm := make([]map[int]string,5)
    // for k,_ := range sm {
    //     sm[k] = make(map[int]string,1)
    //     sm[k][1] = "ok"
    //     fmt.Println(sm[k])
    // }
    // fmt.Println(sm)

    // m := map[int]string{1:"a", 2:"b", 3:"c", 4:"d"}
    // s := make([]int,len(m))
    // i := 0
    // for k,_ := range m {
    //     s[i] = k
    //     i++

    // }
    // sort.Ints(s)
    // fmt.Println(s)

    // m1 := map[int]string{1:"a", 2:"b", 3:"c"}
    // fmt.Println(m1)
    // m2 := make(map[string]int)

    // for k,v:= range m1 {
    //     m2[v] = k
    // }
    // fmt.Println(m2)
    
// }
// func A(a int,v string) (int, string) {

// }
// func A(a, b, c int) int {

// }
// func A() (a, b, c int) {
//     a, b, c = 1, 2, 3
//     return a, b, c
// }

// func A( a ...int) {
//     fmt.Println(a)
// }

// func main() {
//     s1 := []int{1,2,3,4}
//     A(s1)
//     fmt.Println(s1)
// }

// func A(s1 []int) {
//     s1[0] = 5
//     s1[1] = 6
//     s1[2] = 7
//     s1[3] = 8
//     fmt.Println(s1)
// }

// func main() {
//     s1 := 1
//     A(&s1)
//     fmt.Println(s1)
// }

// func A(s1 *int) {
//     *s1 = 2
//     fmt.Println(*s1)
// }

// func main() {
//     s1 := func () {
//         fmt.Println("hello")
//     }
//     s1()
// }

// func closure() {
//     fmt.Println("func A")
// }

// func closure(x,int) func(int) int {
//     retrun func(y int) int {
//     return x + y
//     }
    
// }

// func main() {
//     fmt.Println("a")
//     defer fmt.Println("b")
//     defer fmt.Println("c")
//     }

// func main() {
//     for i := 0; i<3; i++ {
//         defer func () {
//             fmt.Println(i)
//         } ()
//     }
//     }

func main() {
    A()
    B()
    C()
}

func A() {
    fmt.Println("func A")
}

func B() {
    defer func() {
        if err := recover();err != nil {
            fmt.Println("recover in B")
        }
    }()
    panic("panic in B")
}

func C() {
    fmt.Println("func c")
}