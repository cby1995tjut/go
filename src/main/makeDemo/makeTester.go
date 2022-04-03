package main

import "fmt"

func main() {
	//err := justNew()
	newAndMake()

}

func justNew() {
	//panic: runtime error: index out of range [0] with length 0
	s1 := new([]int)
	(*s1)[0] = 100
	fmt.Print(s1)
}

func newAndMake() {
	s1 := new([]int)
	fmt.Println(s1)

	*s1 = make([]int, 10)
	(*s1)[0] = 100
	fmt.Println(s1)
}
