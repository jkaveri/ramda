package main

import (
	"fmt"
	"reflect"

	"github.com/jkaveri/ramda/rstruct"
)

type User struct {
	Name  string
	Age   int
	Email string
}

type Person struct {
	User
	Active bool
}

func main() {
	fmt.Println("=== Struct Reflection Debug ===")
	fmt.Println()

	person := Person{
		User:   User{Name: "Alice", Age: 25, Email: "alice@example.com"},
		Active: true,
	}

	// Debug: Show the struct structure
	fmt.Println("Person Struct Structure:")
	val := reflect.ValueOf(person)
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		fmt.Printf("  Field %d: %s (Type: %s, Anonymous: %t, Kind: %s)\n",
			i, fieldType.Name, field.Type(), fieldType.Anonymous, field.Kind())

		if field.Kind() == reflect.Struct {
			for j := 0; j < field.NumField(); j++ {
				subField := field.Field(j)
				subFieldType := field.Type().Field(j)
				fmt.Printf("    SubField %d: %s (Type: %s, Anonymous: %t)\n",
					j, subFieldType.Name, subField.Type(), subFieldType.Anonymous)
			}
		}
	}
	fmt.Println()

	// Test with standalone functions
	fmt.Println("Field Access Tests:")

	// Test direct field access (only Active is directly accessible)
	fmt.Println("Direct Field Access:")
	active, found := rstruct.Get(person, "Active")
	fmt.Printf("  Get('Active'): %v, found=%t, type=%T\n", active, found, active)

	// Test embedded struct field access using explicit dot notation
	fmt.Println("\nEmbedded Struct Fields (using dot notation):")
	userName, found := rstruct.Get(person, "User.Name")
	fmt.Printf("  Get('User.Name'): %v, found=%t\n", userName, found)

	userAge, found := rstruct.Get(person, "User.Age")
	fmt.Printf("  Get('User.Age'): %v, found=%t\n", userAge, found)

	userEmail, found := rstruct.Get(person, "User.Email")
	fmt.Printf("  Get('User.Email'): %v, found=%t\n", userEmail, found)

	// Test that direct access to embedded fields doesn't work
	fmt.Println("\nDirect Access to Embedded Fields (should fail):")
	name, found := rstruct.Get(person, "Name")
	fmt.Printf("  Get('Name'): %v, found=%t (expected: false)\n", name, found)

	age, found := rstruct.Get(person, "Age")
	fmt.Printf("  Get('Age'): %v, found=%t (expected: false)\n", age, found)

	email, found := rstruct.Get(person, "Email")
	fmt.Printf("  Get('Email'): %v, found=%t (expected: false)\n", email, found)

	// Test Has function
	fmt.Println("\nHas Function Tests:")
	fmt.Printf("  Has('Active'): %t\n", rstruct.Has(person, "Active"))
	fmt.Printf("  Has('User.Name'): %t\n", rstruct.Has(person, "User.Name"))
	fmt.Printf("  Has('Name'): %t (expected: false)\n", rstruct.Has(person, "Name"))

	fmt.Println("\n=== Debug Complete ===")
	fmt.Println("Note: The current implementation requires explicit dot notation")
	fmt.Println("to access embedded struct fields (e.g., 'User.Name' instead of 'Name').")
	fmt.Println("Direct field access only works for fields that are directly")
	fmt.Println("accessible in the struct (like 'Active' in this case).")
}
