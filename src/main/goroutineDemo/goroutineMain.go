package main

import (
	"fmt"
	"runtime"
)

func main() {
	go say("world")
	say("hello")

	fmt.Println("------------")
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c
	z := 0
	fmt.Println(z)
	go give(c)

	z = <-c // receive from c
	fmt.Println(x, y, x+y, z)

	fmt.Println("---------")

	//you can use v, ok := <-c test if a channel is closed
	tmp, ok := <-c
	fmt.Println(tmp)
	fmt.Println(ok)
}

//runtime.Gosched() means let the CPU execute other goroutines, and come back at some point.
func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()

		fmt.Println(s)
	}
}

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total // send total to c
	fmt.Println(total)
	fmt.Println("......")
}

func give(c chan int) {
	c <- 8
	c <- 9
}
