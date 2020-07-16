package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName  string
	age       int
	gender    string
}

// Greeting function (value reciever) for returning some value
// Not for adjusting the value
func (person Person) greet() string {
	return "Hello my name is " + person.firstName + " " + person.lastName + " and I am " + strconv.Itoa(person.age)
}

// Has Birday function (pointer reciever)
// Updating the value
func (person *Person) hasBirthday() {
	person.age++
}

// Get married (pointer reciever)
func (person *Person) getMarried(spouseLastname string) {
	if person.gender == "Male" {
		return
	} else {
		person.lastName = spouseLastname
	}
}

func main() {
	// Common Way
	personOne := Person{firstName: "Huda", lastName: "Prasetyo", age: 17, gender: "Male"}
	// Alternate Way
	// personOne := Person{"Huda", "Prasetyo", 17, "Male"}

	personOne.hasBirthday()
	personOne.getMarried("Doe")
	fmt.Println(personOne.greet())
}
