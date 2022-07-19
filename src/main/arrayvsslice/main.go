package main

import (
	"fmt"
)

func main() {

	var arrstr = [3]string{"1", "2", "3"}

	var sliceStr = []string{"2", "3", "4"}

	ints := make([]int, 2, 2)

	i := append(ints, 3, 4, 5)

	fmt.Print(cap(sliceStr))
	fmt.Print(cap(arrstr))
	fmt.Print(cap(i))
	fmt.Print(cap(ints))
}
