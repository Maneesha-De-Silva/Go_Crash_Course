package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	/*
		firstName string
		lastName  string
		city      string
		gender    string
		age       int
	*/

	firstName, lastName, city, gender string
	age                               int
}

func (p Person) greet() string {
	return "Hello , my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

//hasBrithday method (pointer reciever)
func (p *Person) hasBrithday() {
	p.age++
}

//getMarried (pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "male" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "Female", age: 24}
	person2 := Person{firstName: "Nimantha", lastName: "Smith", city: "Boston", gender: "male", age: 27}

	//Alternative
	//person1 := Person{"Samantha", "Smith", "Boston", "Female", 24}

	/*fmt.Println(person1, person2)
	fmt.Println(person1.firstName)
	person1.age += 5
	fmt.Println(person1.age)
	*/

	person1.hasBrithday()

	person1.getMarried("Williams")
	person2.getMarried("Thompson")

	fmt.Println(person1.greet())
	fmt.Println(person2.greet())
	fmt.Println(person2)

}
