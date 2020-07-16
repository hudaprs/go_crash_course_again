package main

import "fmt"

func main() {
	// Arrays
	numbers := []int{1, 2, 3, 4, 5}

	// With index
	for index, number := range numbers {
		fmt.Printf("%d - ID: %d\n", index, number)
	}

	// Without index
	for _, number := range numbers {
		fmt.Println("Number", number)
	}

	// Maps
	emails := map[string]string{"Bob": "bob@gmail.com", "Mike": "mike@gmail.com"}

	// Loops maps with key
	for key, email := range emails {
		fmt.Printf("Name %s: %s\n", key, email)
	}

	// Loop without key
	for email := range emails {
		fmt.Println("Name", email)
	}
}
