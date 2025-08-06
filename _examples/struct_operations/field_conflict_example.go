package main

import (
	"fmt"

	"github.com/jkaveri/ramda/rstruct"
)

// Test struct with field name conflicts
type Inner struct {
	Name string
	Age  int
}

type Outer struct {
	Name  string // Same name as Inner.Name
	Age   int    // Same name as Inner.Age
	Inner Inner
}

// Test struct with embedded field name conflicts
type Embedded struct {
	Name string
	Age  int
}

type Container struct {
	Embedded        // Embedded struct
	Name     string // Same name as Embedded.Name
	Age      int    // Same name as Embedded.Age
}

func main() {
	fmt.Println("=== Field Name Conflict Examples ===")
	fmt.Println()

	// Test 1: Named struct field conflicts
	fmt.Println("1. Named Struct Field Conflicts:")
	fmt.Println("   Outer{Name: 'OuterName', Age: 30, Inner: Inner{Name: 'InnerName', Age: 25}}")

	outer := Outer{
		Name:  "OuterName",
		Age:   30,
		Inner: Inner{Name: "InnerName", Age: 25},
	}

	// Test outer level fields
	name, found := rstruct.Get(outer, "Name")
	fmt.Printf("   Get('Name'): %v, found=%t\n", name, found)

	age, found := rstruct.Get(outer, "Age")
	fmt.Printf("   Get('Age'): %v, found=%t\n", age, found)

	// Test inner level fields using dot notation
	innerName, found := rstruct.Get(outer, "Inner.Name")
	fmt.Printf("   Get('Inner.Name'): %v, found=%t\n", innerName, found)

	innerAge, found := rstruct.Get(outer, "Inner.Age")
	fmt.Printf("   Get('Inner.Age'): %v, found=%t\n", innerAge, found)
	fmt.Println()

	// Test 2: Embedded struct field conflicts
	fmt.Println("2. Embedded Struct Field Conflicts:")
	fmt.Println("   Container{Embedded: Embedded{Name: 'EmbeddedName', Age: 25}, Name: 'ContainerName', Age: 30}")

	container := Container{
		Embedded: Embedded{Name: "EmbeddedName", Age: 25},
		Name:     "ContainerName",
		Age:      30,
	}

	// Test field access (should show which one takes precedence)
	name2, found := rstruct.Get(container, "Name")
	fmt.Printf("   Get('Name'): %v, found=%t (which one?)\n", name2, found)

	age2, found := rstruct.Get(container, "Age")
	fmt.Printf("   Get('Age'): %v, found=%t (which one?)\n", age2, found)
	fmt.Println()

	// Test 3: Show field precedence behavior
	fmt.Println("3. Field Precedence Behavior:")

	// Using standalone functions
	standaloneName, found := rstruct.Get(container, "Name")
	fmt.Printf("   Get('Name'): %v, found=%t\n", standaloneName, found)

	standaloneAge, found := rstruct.Get(container, "Age")
	fmt.Printf("   Get('Age'): %v, found=%t\n", standaloneAge, found)

	// Test embedded struct fields explicitly
	embeddedName, found := rstruct.Get(container, "Embedded.Name")
	fmt.Printf("   Get('Embedded.Name'): %v, found=%t\n", embeddedName, found)

	embeddedAge, found := rstruct.Get(container, "Embedded.Age")
	fmt.Printf("   Get('Embedded.Age'): %v, found=%t\n", embeddedAge, found)
	fmt.Println()

	// Test 4: Set operations with field conflicts
	fmt.Println("4. Set Operations with Field Conflicts:")

	containerPtr := &container

	// Set the top-level Name field
	err := rstruct.Set(containerPtr, "Name", "NewContainerName")
	if err == nil {
		fmt.Printf("   Set('Name'): Success\n")
	} else {
		fmt.Printf("   Set('Name'): Error - %v\n", err)
	}

	// Set the embedded Name field
	err = rstruct.Set(containerPtr, "Embedded.Name", "NewEmbeddedName")
	if err == nil {
		fmt.Printf("   Set('Embedded.Name'): Success\n")
	} else {
		fmt.Printf("   Set('Embedded.Name'): Error - %v\n", err)
	}

	// Verify the changes
	updatedName, _ := rstruct.Get(*containerPtr, "Name")
	updatedEmbeddedName, _ := rstruct.Get(*containerPtr, "Embedded.Name")
	fmt.Printf("   After Set - Top Name: %v, Embedded Name: %v\n", updatedName, updatedEmbeddedName)
	fmt.Println()

	// Test 5: Has function with field conflicts
	fmt.Println("5. Has Function with Field Conflicts:")

	fmt.Printf("   Has('Name'): %t\n", rstruct.Has(container, "Name"))
	fmt.Printf("   Has('Age'): %t\n", rstruct.Has(container, "Age"))
	fmt.Printf("   Has('Embedded.Name'): %t\n", rstruct.Has(container, "Embedded.Name"))
	fmt.Printf("   Has('Embedded.Age'): %t\n", rstruct.Has(container, "Embedded.Age"))
	fmt.Printf("   Has('NonExistent'): %t\n", rstruct.Has(container, "NonExistent"))
	fmt.Println()

	fmt.Println("=== Field Conflict Examples Complete ===")
	fmt.Println()
	fmt.Println("Note: When there are field name conflicts between embedded structs")
	fmt.Println("and the containing struct, the implementation typically gives")
	fmt.Println("precedence to the top-level fields. Use explicit dot notation")
	fmt.Println("to access embedded struct fields.")
}
