package main

import "fmt"

func main() {
	//var fruitArr [2]string

	//fruitArr[0] = "Apple"
	//fruitArr[1] = "Orange"
	fruitArr := [3]string{"Apple", "Orange", "Grape"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])

	fruitSlice := []string{"Apple", "Orange", "Grape"}
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:2])
}
