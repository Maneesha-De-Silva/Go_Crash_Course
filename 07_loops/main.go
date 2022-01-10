package main

import "fmt"

func main() {
	// Log method
	/*i := 1
	for i <= 10 {
		fmt.Println(i)
		//i = i+1
		i++
	}

	// short method
	for i := 1; i <= 10; i++ {
		fmt.Printf("Number %d\n", i)
	}

	//FizzBuzz
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}*/

	// Define map
	//emails := make(map[string]string)

	// Assign kv
	//emails["Bob"] = "bob@gmail.com"
	//emails["Sharon"] = "sharon@gmail.com"
	//emails["Mike"] = "mike@gmail.com"

	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	//Delete from maps
	delete(emails, "Bob")
	fmt.Println(emails)
}
