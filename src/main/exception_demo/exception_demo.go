package main

import (
	"fmt"
)

func main() {
	test()

	fmt.Println("hello, my world")
}

func test() {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	i := 1
	j := 0

	c := i / j

	fmt.Println(&c)
}
