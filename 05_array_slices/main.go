package main

import "fmt"

func main() {
	// Array
	// var fruits [2]string
	// fruits[0] = "Mango"
	// fruits[1] = "Apple"

	// Slices / Shorthand
	fruits := []string{"Apple", "Mango", "Grape", "Cherry"}

	fmt.Println(fruits)
	// Count the array items
	fmt.Println(len(fruits))
	// Get item with reange
	fmt.Println(fruits[0:2])
}
