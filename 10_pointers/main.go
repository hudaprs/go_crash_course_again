package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a)
	fmt.Printf("%T\n", b)

	// Use * to read the pointers / memory adderss value
	fmt.Println(*b)

	// Change the pointer value
	*b = 10
	fmt.Println(a)
}
