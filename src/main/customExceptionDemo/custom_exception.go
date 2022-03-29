package main

import (
	"errors"
	"fmt"
)

func main() {
	err := readConf("aaa")
	if err == nil {
		fmt.Print("Continue....")
	} else {
		panic(err)
	}

	fmt.Println("Not disturbed")
}

func readConf(name string) (err error) {

	if "conf" == name {
		return nil
	} else {
		return errors.New("not found")
	}
}
