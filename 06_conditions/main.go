package main

import (
	"fmt"
)

func main() {
	x := 15
	y := 10

	//if else
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	// else if
	color := "green"

	if color == "red" {
		fmt.Println("Color is Red")
	} else if color == "blue" {
		fmt.Println("Color is Blue")
	} else {
		fmt.Println("Color is not red or blue")
	}
}
