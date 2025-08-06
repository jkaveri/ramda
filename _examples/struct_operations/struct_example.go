package main

import (
	"fmt"

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
	fmt.Println("Struct Operations Examples")
	fmt.Println("=========================")
	fmt.Println()

	// 1. Get - Retrieve field values
	fmt.Println("1. Get - Retrieve field values:")
	user := User{Name: "Alice", Age: 25, Email: "alice@example.com"}
	fmt.Printf("  Original: %+v\n", user)

	value, found := rstruct.Get(user, "Name")
	if found {
		fmt.Printf("  Get 'Name': %v\n", value)
	}

	value, found = rstruct.Get(user, "Age")
	if found {
		fmt.Printf("  Get 'Age': %v\n", value)
	}

	// Case-sensitive
	value, found = rstruct.Get(user, "Email")
	if found {
		fmt.Printf("  Get 'Email': %v\n", value)
	}

	// Non-existing field
	_, found = rstruct.Get(user, "Phone")
	if !found {
		fmt.Println("  Get 'Phone': field not found (as expected)")
	}
	fmt.Println()

	// 2. GetOrDefault - Safe field access with defaults
	fmt.Println("2. GetOrDefault - Safe field access with defaults:")
	value = rstruct.GetOrDefault(user, "Age", 0)
	fmt.Printf("  GetOrDefault 'Age' with default 0: %v\n", value)

	value = rstruct.GetOrDefault(user, "Phone", "unknown")
	fmt.Printf("  GetOrDefault 'Phone' with default 'unknown': %v\n", value)
	fmt.Println()

	// 3. Set - Update field values
	fmt.Println("3. Set - Update field values:")
	userPtr := &user
	err := rstruct.Set(userPtr, "Age", 30)
	if err == nil {
		fmt.Printf("  Set 'Age' to 30: %+v\n", *userPtr)
	} else {
		fmt.Printf("  Error setting 'Age': %v\n", err)
	}

	err = rstruct.Set(userPtr, "Email", "new@example.com")
	if err == nil {
		fmt.Printf("  Set 'Email' to 'new@example.com': %+v\n", *userPtr)
	} else {
		fmt.Printf("  Error setting 'Email': %v\n", err)
	}

	// Case-sensitive
	err = rstruct.Set(userPtr, "Name", "Bob")
	if err == nil {
		fmt.Printf("  Set 'Name' to 'Bob': %+v\n", *userPtr)
	} else {
		fmt.Printf("  Error setting 'Name': %v\n", err)
	}

	// Non-existing field
	err = rstruct.Set(userPtr, "Phone", "123-456-7890")
	if err != nil {
		fmt.Printf("  Set 'Phone': %v (as expected)\n", err)
	}
	fmt.Println()

	// 4. Has - Check field existence
	fmt.Println("4. Has - Check field existence:")
	fmt.Printf("  Has 'Name': %t\n", rstruct.Has(user, "Name"))
	fmt.Printf("  Has 'Age': %t\n", rstruct.Has(user, "Age"))
	fmt.Printf("  Has 'Phone': %t\n", rstruct.Has(user, "Phone"))
	fmt.Printf("  Has 'NAME' (case-sensitive): %t\n", rstruct.Has(user, "NAME"))
	fmt.Println()

	// 5. Fields - Get all field names
	fmt.Println("5. Fields - Get all field names:")
	fields := rstruct.Fields(user)
	fmt.Printf("  Fields: %v\n", fields)
	fmt.Println()



	// 6. ToMap - Convert struct to map
	fmt.Println("6. ToMap - Convert struct to map:")
	result := rstruct.ToMap(user)
	fmt.Printf("  ToMap: %v\n", result)
	fmt.Println()

	// 7. FromMap - Create struct from map
	fmt.Println("7. FromMap - Create struct from map:")
	data := map[string]interface{}{
		"Name":  "Charlie",
		"Age":   35,
		"Email": "charlie@example.com",
	}

	userFromMap, success := rstruct.FromMap[User](data)
	if success {
		fmt.Printf("  FromMap: %+v\n", userFromMap)
	}

	// Test with extra fields
	data["Extra"] = "ignored"
	userFromMap, success = rstruct.FromMap[User](data)
	if success {
		fmt.Printf("  FromMap with extra fields: %+v\n", userFromMap)
	}
	fmt.Println()

	// 8. Pick - Select specific fields
	fmt.Println("8. Pick - Select specific fields:")
	picked := rstruct.Pick(user, []string{"Name", "Age"})
	fmt.Printf("  Pick ['Name', 'Age']: %+v\n", picked)

	// Case-sensitive
	picked = rstruct.Pick(user, []string{"Name", "Email"})
	fmt.Printf("  Pick ['Name', 'Email']: %+v\n", picked)

	// Non-existing fields
	picked = rstruct.Pick(user, []string{"Name", "Phone"})
	fmt.Printf("  Pick ['Name', 'Phone'] (non-existing ignored): %+v\n", picked)
	fmt.Println()

	// 9. Omit - Exclude specific fields
	fmt.Println("9. Omit - Exclude specific fields:")
	omitted := rstruct.Omit(user, []string{"Email"})
	fmt.Printf("  Omit ['Email']: %+v\n", omitted)

	// Case-sensitive
	omitted = rstruct.Omit(user, []string{"Email", "Age"})
	fmt.Printf("  Omit ['Email', 'Age']: %+v\n", omitted)

	// Non-existing fields
	omitted = rstruct.Omit(user, []string{"Phone"})
	fmt.Printf("  Omit ['Phone'] (non-existing ignored): %+v\n", omitted)
	fmt.Println()

	// 10. Merge - Combine multiple structs
	fmt.Println("10. Merge - Combine multiple structs:")
	user1 := User{Name: "Alice", Age: 25, Email: "alice@example.com"}
	user2 := User{Name: "Bob", Age: 30, Email: ""}
	user3 := User{Name: "", Age: 0, Email: "charlie@example.com"}

	merged := rstruct.Merge(user1, user2)
	fmt.Printf("  Merge user1 + user2: %+v\n", merged)

	merged = rstruct.Merge(user1, user2, user3)
	fmt.Printf("  Merge user1 + user2 + user3: %+v\n", merged)
	fmt.Println()

	// 11. Clone - Create deep copy
	fmt.Println("11. Clone - Create deep copy:")
	cloned := rstruct.Clone(user)
	fmt.Printf("  Cloned: %+v\n", cloned)

	// Verify it's a different instance
	if &cloned != &user {
		fmt.Println("  ✓ Cloned struct is a different instance")
	}
	fmt.Println()

	// 12. Nested struct operations
	fmt.Println("12. Nested struct operations:")
	person := Person{
		User:    User{Name: "Alice", Age: 25, Email: "alice@example.com"},
		Address: Address{Street: "123 Main St", City: "New York", Country: "USA"},
		Active:  true,
	}

	fmt.Printf("  Original person: %+v\n", person)

	// Get from nested struct
	value, found = rstruct.Get(person, "Name")
	if found {
		fmt.Printf("  Get 'Name' from nested struct: %v\n", value)
	}

	// Set in nested struct
	personPtr := &person
	err = rstruct.Set(personPtr, "Active", false)
	if err == nil {
		fmt.Printf("  Set 'Active' in nested struct: %+v\n", *personPtr)
	} else {
		fmt.Printf("  Error setting 'Active': %v\n", err)
	}

	// Fields of nested struct
	personFields := rstruct.Fields(person)
	fmt.Printf("  Fields of nested struct: %v\n", personFields)

	// ToMap of nested struct
	personMap := rstruct.ToMap(person)
	fmt.Printf("  ToMap of nested struct: %v\n", personMap)
	fmt.Println()

	// 14. Function composition with struct operations
	fmt.Println("14. Function composition with struct operations:")

	// Create a user, update it, pick specific fields, then convert to map
	originalUser := User{Name: "David", Age: 28, Email: "david@example.com"}

	// Chain operations: Set -> Pick -> ToMap
	originalUserPtr := &originalUser
	err = rstruct.Set(originalUserPtr, "Age", 29)
	if err == nil {
		pickedUser := rstruct.Pick(*originalUserPtr, []string{"Name", "Age"})
		finalMap := rstruct.ToMap(pickedUser)

		fmt.Printf("  Original: %+v\n", originalUser)
		fmt.Printf("  After Set age=29: %+v\n", *originalUserPtr)
		fmt.Printf("  After Pick ['Name', 'Age']: %+v\n", pickedUser)
		fmt.Printf("  Final ToMap: %v\n", finalMap)
	}

	// Another composition: FromMap -> Set -> Omit
	data2 := map[string]interface{}{
		"Name":  "Eve",
		"Age":   32,
		"Email": "eve@example.com",
	}

	userFromData, _ := rstruct.FromMap[User](data2)
	userFromDataPtr := &userFromData
	err = rstruct.Set(userFromDataPtr, "Email", "eve.new@example.com")
	if err == nil {
		userWithoutEmail := rstruct.Omit(*userFromDataPtr, []string{"Email"})

		fmt.Printf("  FromMap: %+v\n", userFromData)
		fmt.Printf("  After Set email: %+v\n", *userFromDataPtr)
		fmt.Printf("  After Omit ['Email']: %+v\n", userWithoutEmail)
	}
	fmt.Println()

	// 15. Error handling examples
	fmt.Println("15. Error handling examples:")

	// Test with non-struct types
	str := "hello"
	_, found = rstruct.Get(str, "field")
	if !found {
		fmt.Println("  ✓ Get on string correctly returns not found")
	}

	if !rstruct.Has(str, "field") {
		fmt.Println("  ✓ Has on string correctly returns false")
	}

	// Test type mismatch in Set
	err = rstruct.Set(userPtr, "Age", "thirty")
	if err != nil {
		fmt.Printf("  ✓ Set with type mismatch correctly fails: %v\n", err)
	}

	// Test non-existing field in Set
	err = rstruct.Set(userPtr, "Phone", "123-456-7890")
	if err != nil {
		fmt.Printf("  ✓ Set with non-existing field correctly fails: %v\n", err)
	}
}
