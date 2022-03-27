package main

import "fmt"

func main() {
	num1 := 100
	fmt.Printf("类型%T \n", num1)
	fmt.Printf("num1的值=%v \n", num1)
	fmt.Printf("num1的地址%v \n", &num1)

	num2 := new(int)
	*num2 = 100
	fmt.Println("------------split--------------")

	fmt.Printf("类型%T \n", num2)
	fmt.Printf("num2的值%v \n", num2)
	fmt.Printf("num2的地址%v \n", &num2)
}
