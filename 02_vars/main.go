package main

import "fmt"

func main() {
	// var name = "Huda"
	// const name := "Huda"
	// weight := 1.9

	const name = "Huda"
	var age = "17"
	var isCool = false

	// Cast the type
	isCool = true

	fmt.Println(name, age, isCool)
	fmt.Printf("%T\n", isCool)
	// fmt.Printf("%T", weight)
}
