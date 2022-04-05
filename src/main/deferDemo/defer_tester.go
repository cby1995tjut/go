package main

//延迟处理当前函数内的defer函数,
//相当于把defer 引用的方法压入栈中，等函数调用完
import (
	"fmt"
)

func sum(n1 int, n2 int) int {
	defer fmt.Println(n1)
	defer fmt.Println(n2)

	res := n1 + n2
	fmt.Println(res)
	return res
}

func main() {
	res := sum(10, 20)
	fmt.Println("result", res)
}
