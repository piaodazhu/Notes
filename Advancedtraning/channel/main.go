package main

import (
	"fmt"
	"time"
)

func main() {
	var a chan int
	b := make(chan int)

	go func() {
		for x := range b {
			fmt.Println(x)
		}
		x, ok := <-b
		fmt.Println(x, ok)
		fmt.Println("go exit")
	}()
	// time.Sleep(time.Second)
	b <- 0
	b <- 1
	b <- 2
	close(b)
	// b <- 3
	time.Sleep(time.Second)
	fmt.Printf("%T, %v, %#v\n", a, a, a)
	fmt.Printf("%T, %v, %#v\n", b, b, b)
}
