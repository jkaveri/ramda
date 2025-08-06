package main

import (
	"fmt"

	"github.com/jkaveri/ramda/rstruct"
)

type Address struct {
	Street  string
	City    string
	Country string
}

type User struct {
	Name  string
	Age   int
	Email string
}

type Person struct {
	User
	Address Address
	Active  bool
}

func main() {
	fmt.Println("=== Nested Field Access Test ===")
	fmt.Println()

	person := Person{
		User:    User{Name: "Alice", Age: 25, Email: "alice@example.com"},
		Address: Address{Street: "123 Main St", City: "New York", Country: "USA"},
		Active:  true,
	}

	// Test nested field access - first access
	fmt.Println("First Access:")
	city1, found1 := rstruct.Get(person, "Address.City")
	fmt.Printf("  Get('Address.City'): %v, found=%t\n", city1, found1)

	street1, found1 := rstruct.Get(person, "Address.Street")
	fmt.Printf("  Get('Address.Street'): %v, found=%t\n", street1, found1)

	country1, found1 := rstruct.Get(person, "Address.Country")
	fmt.Printf("  Get('Address.Country'): %v, found=%t\n", country1, found1)

	// Test nested field access - second access
	fmt.Println("\nSecond Access:")
	city2, found2 := rstruct.Get(person, "Address.City")
	fmt.Printf("  Get('Address.City'): %v, found=%t\n", city2, found2)

	street2, found2 := rstruct.Get(person, "Address.Street")
	fmt.Printf("  Get('Address.Street'): %v, found=%t\n", street2, found2)

	country2, found2 := rstruct.Get(person, "Address.Country")
	fmt.Printf("  Get('Address.Country'): %v, found=%t\n", country2, found2)

	// Test embedded field access
	fmt.Println("\nEmbedded Field Access:")
	name1, found1 := rstruct.Get(person, "Name")
	fmt.Printf("  Get('Name'): %v, found=%t\n", name1, found1)

	age1, found1 := rstruct.Get(person, "Age")
	fmt.Printf("  Get('Age'): %v, found=%t\n", age1, found1)

	email1, found1 := rstruct.Get(person, "Email")
	fmt.Printf("  Get('Email'): %v, found=%t\n", email1, found1)

	// Test nested field setting
	fmt.Println("\nNested Field Setting:")
	personPtr := &person
	err := rstruct.Set(personPtr, "Address.City", "Los Angeles")
	if err == nil {
		fmt.Printf("  Set('Address.City', 'Los Angeles'): Success\n")
		fmt.Printf("  Result - City: %s\n", personPtr.Address.City)
	} else {
		fmt.Printf("  Set('Address.City', 'Los Angeles'): Error - %v\n", err)
	}

	// Test field existence
	fmt.Println("\nField Existence:")
	fmt.Printf("  Has('Address.City'): %t\n", rstruct.Has(person, "Address.City"))
	fmt.Printf("  Has('Address.Street'): %t\n", rstruct.Has(person, "Address.Street"))
	fmt.Printf("  Has('Address.Country'): %t\n", rstruct.Has(person, "Address.Country"))
	fmt.Printf("  Has('Address.Phone'): %t\n", rstruct.Has(person, "Address.Phone"))

	// Test multiple accesses
	fmt.Println("\nMultiple Accesses:")
	for i := 0; i < 5; i++ {
		city, found := rstruct.Get(person, "Address.City")
		fmt.Printf("  Access %d - Get('Address.City'): %v, found=%t\n", i+1, city, found)
	}

	// Test deep nested access
	fmt.Println("\nDeep Nested Access:")
	deepCity, found := rstruct.Get(person, "Address.City")
	fmt.Printf("  Get('Address.City'): %v, found=%t\n", deepCity, found)

	// Test setting deep nested field
	err = rstruct.Set(personPtr, "Address.Country", "Canada")
	if err == nil {
		fmt.Printf("  Set('Address.Country', 'Canada'): Success\n")
		fmt.Printf("  Result - Country: %s\n", personPtr.Address.Country)
	} else {
		fmt.Printf("  Set('Address.Country', 'Canada'): Error - %v\n", err)
	}

	fmt.Println("\n=== Nested Field Access Test Complete ===")
	fmt.Println("Note: The current implementation supports nested field access")
	fmt.Println("using dot notation (e.g., 'Address.City') and can set")
	fmt.Println("nested fields when using pointers to the struct.")
}
