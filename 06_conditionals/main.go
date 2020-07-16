package main

import "fmt"

func main() {
	x := 5
	y := 10

	// If Statement
	if x == y {
		fmt.Printf("%d is equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is NOT equal to %d\n", x, y)
	}

	// Else If Statement
	color := "red"
	if color == "red" {
		fmt.Println("Color is red")
	} else if color == "blue" {
		fmt.Println("Color is blue")
	} else {
		fmt.Println("Color is NOT blue or red")
	}

	// Switch
	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Color is blue")
	default:
		fmt.Println("Color is NOT blue or red")
	}
}
