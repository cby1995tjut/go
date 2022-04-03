package main

import "fmt"

const (
	x = iota
	y = iota
	z = iota
	w = iota
)

func myIota() {
	fmt.Print(x, y, z, w)
}
