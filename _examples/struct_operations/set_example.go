package main

import (
	"fmt"
	"log"

	"github.com/jkaveri/ramda/rstruct"
)

type User struct {
	Name  string
	Age   int
	Email string
}

type Address struct {
	Street  string
	City    string
	Country string
}

type Person struct {
	User
	Address Address
	Active  bool
}

func main() {
	fmt.Println("=== Set Function Example (Returns error) ===")

	// Example 1: Setting a simple field
	user := &User{Name: "Alice", Age: 25, Email: "alice@example.com"}
	fmt.Printf("Before: %+v\n", user)

	err := rstruct.Set(user, "Age", 30)
	if err != nil {
		log.Printf("Error setting Age: %v", err)
	} else {
		fmt.Printf("After setting Age to 30: %+v\n", user)
	}

	// Example 2: Setting nested field
	person := &Person{
		User:    User{Name: "Bob", Age: 30, Email: "bob@example.com"},
		Address: Address{Street: "123 Main St", City: "New York", Country: "USA"},
		Active:  true,
	}
	fmt.Printf("\nBefore: %+v\n", person)

	err = rstruct.Set(person, "Address.City", "Los Angeles")
	if err != nil {
		log.Printf("Error setting Address.City: %v", err)
	} else {
		fmt.Printf("After setting Address.City to 'Los Angeles': %+v\n", person)
	}

	// Example 3: Setting embedded struct field
	err = rstruct.Set(person, "User.Email", "bob.new@example.com")
	if err != nil {
		log.Printf("Error setting User.Email: %v", err)
	} else {
		fmt.Printf("After setting User.Email: %+v\n", person)
	}

	// Example 4: Error cases
	fmt.Println("\n=== Error Examples ===")

	// Non-existing field
	err = rstruct.Set(user, "Phone", "123-456-7890")
	if err != nil {
		fmt.Printf("Expected error for non-existing field: %v\n", err)
	}

	// Type mismatch
	err = rstruct.Set(user, "Age", "thirty")
	if err != nil {
		fmt.Printf("Expected error for type mismatch: %v\n", err)
	}

	// Non-struct type
	var str string = "hello"
	err = rstruct.Set(&str, "field", "value")
	if err != nil {
		fmt.Printf("Expected error for non-struct type: %v\n", err)
	}
}
