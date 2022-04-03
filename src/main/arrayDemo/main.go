package main

import "fmt"

func main() {
	var hens [2]float64
	hens[0] = 1.0
	hens[1] = 2.0
	total := 0.0

	for i := 0; i < len(hens); i++ {
		total += hens[i]
	}
	fmt.Print(total)

	fmt.Println("--------------------------")

	myIota()
}
