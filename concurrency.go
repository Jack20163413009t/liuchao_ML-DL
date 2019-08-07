package main

import (
  "fmt"
  // "runtime"
  // "sync"
  )

// func main() {
// 	c := make(chan bool)
// 	go func() {
// 		fmt.Println("go go go")
// 		c <- true
// 		close(c)
// 	} ()
// 	for v := range c {
// 		fmt.Println(v)
// 	}
// }


// func main() {
// 	runtime.GOMAXPROCS(runtime.NumCPU())
// 	c := make(chan bool, 10)
// 	for i:= 0; i <10; i++ {
// 		go Go(c, i)
// 	}

// 	for i := 0; i <10; i++ {
// 		<- c
// 	}
// }


// func Go(c chan bool, index int) {
// 	a := 1
// 	for i := 0; i < 10000000; i++ {
// 		a += i 
// 	}
// 	fmt.Println(index, a)

// 	c <- true
// }


// func main() {
// 	runtime.GOMAXPROCS(runtime.NumCPU())
// 	wg := sync.WaitGroup{}
// 	wg.Add(10)
// 	for i:= 0; i <10; i++ {
// 		go Go(&wg, i)
// 	}
// 	wg.Wait()
// }


// func Go(wg *sync.WaitGroup, index int) {
// 	a := 1
// 	for i := 0; i < 10000000; i++ {
// 		a += i 
// 	}
// 	fmt.Println(index, a)

// 	wg.Done()
// }


func main() {
	c1, c2 := make(chan int), make(chan string)
	o := make(chan bool, 2)
	go func() {
		a, b := false,false
		for {
			select {
			case v,ok := <-c1:
				if !ok {
					if !a {
						a = true
					o <- true
				}
					break
				}
				fmt.Println("c1",v)
			case v,ok := <-c2:
				if !ok {
					if !b {
						b =true
					o <- true
				}
					break
				}
				fmt.Println("c2",v)
			}
		}
	}()

	c1 <- 1
	c2 <- "hi"
	c1 <- 3
	c2 <- "hello"

	close(c1)
	close(c2)

	for i := 0; i<2; i++ {
		<- o
	}
}


