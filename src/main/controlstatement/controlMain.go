package main

import (
	"fmt"
	"time"
)

func main() {
	gotoFunc()
}

func gotoFunc() {
	i := 0
Here:
	fmt.Print(i)
	i++
	time.Sleep(1000 * time.Millisecond)
	goto Here
}
