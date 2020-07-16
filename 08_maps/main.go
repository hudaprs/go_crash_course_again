package main

import "fmt"

func main() {
	// Long Method
	// emails := make(map[string]string)

	// emails["bob"] = "bob@gmail.com"
	// emails["mike"] = "mike@gmail.com"

	// Shorthand
	emails := map[string]string{"bob": "bob@gmail.com", "mike": "mike@gmail.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))

	// Adding new values
	emails["sharon"] = "sharon@gmail.com"
	fmt.Println(emails)

	// Delete some value
	delete(emails, "bob")

	fmt.Println("Deleted", emails)
}
