package main

import (
	"fmt"
)

func main() {

	msg := hello()

	fmt.Println(msg)
}

func hello() string {
	return "Hello world"
}
