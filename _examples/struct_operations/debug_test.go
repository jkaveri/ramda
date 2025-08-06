package main

import (
	"fmt"

	"github.com/jkaveri/ramda/rstruct"
)

type User struct {
	Name string
	Age  int
}

type Person struct {
	User
	Active bool
}

func main() {
	person := Person{
		User:   User{Name: "Alice", Age: 25},
		Active: true,
	}

	fmt.Printf("Person: %+v\n", person)

	// Test direct field access
	name, found := rstruct.Get(person, "Name")
	fmt.Printf("Get('Name'): %v, found=%t\n", name, found)

	age, found := rstruct.Get(person, "Age")
	fmt.Printf("Get('Age'): %v, found=%t\n", age, found)

	active, found := rstruct.Get(person, "Active")
	fmt.Printf("Get('Active'): %v, found=%t\n", active, found)

	// Test nested field access
	userName, found := rstruct.Get(person, "User.Name")
	fmt.Printf("Get('User.Name'): %v, found=%t\n", userName, found)

	userAge, found := rstruct.Get(person, "User.Age")
	fmt.Printf("Get('User.Age'): %v, found=%t\n", userAge, found)

	// Test Has function
	fmt.Printf("Has('Name'): %t\n", rstruct.Has(person, "Name"))
	fmt.Printf("Has('Age'): %t\n", rstruct.Has(person, "Age"))
	fmt.Printf("Has('User.Name'): %t\n", rstruct.Has(person, "User.Name"))
	fmt.Printf("Has('User.Age'): %t\n", rstruct.Has(person, "User.Age"))
}
