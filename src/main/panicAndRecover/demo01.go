package main

import (
	"fmt"
)

func func1() {
	a := "a"
	if a == "a" {
		panic("Nothing need to do")
	}

	fmt.Print("HAHAHAHA")
}
func main() {
	throwsPanic()
	fmt.Println("CCCCCC")
}

// defer + recover 相当与try catch
func throwsPanic() (b bool) {
	defer func() {
		if x := recover(); x != nil {
			b = true
		}
	}()
	func1()
	fmt.Println("DIDIDIDIDI") // if f causes panic, it will recover
	return
}

func init() {
	fmt.Println("Im init")
}

func init() {
	fmt.Println("Im init2")
}
