package main

import (
	"fmt"

	"github.com/jkaveri/ramda/rstruct"
)

// Example structs for demonstration
type Address struct {
	Street  string
	City    string
	Country string
}

type Company struct {
	Name    string
	Address Address
}

type Employee struct {
	Name    string
	Age     int
	Email   string
	Company Company
	Active  bool
}

type Department struct {
	Name      string
	Manager   Employee
	Employees []Employee
}

func main() {
	fmt.Println("=== Nested Field Get/Set Examples ===\n")

	// Create a sample employee
	employee := Employee{
		Name:  "Alice Johnson",
		Age:   30,
		Email: "alice@company.com",
		Company: Company{
			Name: "Tech Corp",
			Address: Address{
				Street:  "123 Tech Street",
				City:    "San Francisco",
				Country: "USA",
			},
		},
		Active: true,
	}

	fmt.Println("Original Employee:")
	printEmployee(employee)

	// Example 1: Get nested field values
	fmt.Println("\n--- Getting Nested Fields ---")

	// Get company name
	if companyName, found := rstruct.Get(employee, "Company.Name"); found {
		fmt.Printf("Company Name: %s\n", companyName)
	}

	// Get company address city
	if city, found := rstruct.Get(employee, "Company.Address.City"); found {
		fmt.Printf("Company City: %s\n", city)
	}

	// Get employee email
	if email, found := rstruct.Get(employee, "Email"); found {
		fmt.Printf("Employee Email: %s\n", email)
	}

	// Example 2: Set nested field values
	fmt.Println("\n--- Setting Nested Fields ---")

	// Update company city
	employeePtr := &employee
	err := rstruct.Set(employeePtr, "Company.Address.City", "Seattle")
	if err == nil {
		fmt.Println("Updated company city to Seattle:")
		printEmployee(*employeePtr)
	} else {
		fmt.Printf("Error updating company city: %v\n", err)
	}

	// Update employee age
	err = rstruct.Set(employeePtr, "Age", 31)
	if err == nil {
		fmt.Println("Updated employee age to 31:")
		printEmployee(*employeePtr)
	} else {
		fmt.Printf("Error updating employee age: %v\n", err)
	}

	// Update company name
	err = rstruct.Set(employeePtr, "Company.Name", "New Tech Corp")
	if err == nil {
		fmt.Println("Updated company name:")
		printEmployee(*employeePtr)
	} else {
		fmt.Printf("Error updating company name: %v\n", err)
	}

	// Example 3: Check if nested fields exist
	fmt.Println("\n--- Checking Nested Fields ---")

	fieldsToCheck := []string{
		"Company.Address.City",
		"Company.Address.Phone", // This doesn't exist
		"Email",
		"Company.Name",
		"Invalid.Field", // This doesn't exist
	}

	for _, field := range fieldsToCheck {
		if rstruct.Has(employee, field) {
			fmt.Printf("✓ Field '%s' exists\n", field)
		} else {
			fmt.Printf("✗ Field '%s' does not exist\n", field)
		}
	}

	// Example 4: Deep nested structures
	fmt.Println("\n--- Deep Nested Structures ---")

	department := Department{
		Name: "Engineering",
		Manager: Employee{
			Name:  "Bob Smith",
			Age:   35,
			Email: "bob@company.com",
			Company: Company{
				Name: "Tech Corp",
				Address: Address{
					Street:  "456 Engineering Ave",
					City:    "San Francisco",
					Country: "USA",
				},
			},
			Active: true,
		},
		Employees: []Employee{
			{
				Name:  "Charlie Brown",
				Age:   28,
				Email: "charlie@company.com",
				Company: Company{
					Name: "Tech Corp",
					Address: Address{
						Street:  "456 Engineering Ave",
						City:    "San Francisco",
						Country: "USA",
					},
				},
				Active: true,
			},
		},
	}

	fmt.Println("Department Manager Info:")
	if managerName, found := rstruct.Get(department, "Manager.Name"); found {
		fmt.Printf("Manager Name: %s\n", managerName)
	}

	if managerCompany, found := rstruct.Get(department, "Manager.Company.Name"); found {
		fmt.Printf("Manager Company: %s\n", managerCompany)
	}

	if managerCity, found := rstruct.Get(department, "Manager.Company.Address.City"); found {
		fmt.Printf("Manager Company City: %s\n", managerCity)
	}

	// Update manager's company city
	departmentPtr := &department
	err = rstruct.Set(departmentPtr, "Manager.Company.Address.City", "New York")
	if err == nil {
		fmt.Println("\nUpdated manager's company city to New York:")
		if city, found := rstruct.Get(*departmentPtr, "Manager.Company.Address.City"); found {
			fmt.Printf("New Manager Company City: %s\n", city)
		}
	} else {
		fmt.Printf("Error updating manager's company city: %v\n", err)
	}

	// Example 5: Case sensitive field access
	fmt.Println("\n--- Case Sensitive Access ---")

	if city, found := rstruct.Get(employee, "Company.Address.City"); found {
		fmt.Printf("Company City: %s\n", city)
	}

	err = rstruct.Set(employeePtr, "Company.Address.Country", "Canada")
	if err == nil {
		fmt.Println("Updated country:")
		if country, found := rstruct.Get(*employeePtr, "Company.Address.Country"); found {
			fmt.Printf("New Country: %s\n", country)
		}
	} else {
		fmt.Printf("Error updating country: %v\n", err)
	}

	// Example 6: Error handling
	fmt.Println("\n--- Error Handling ---")

	// Try to get non-existent field
	if _, found := rstruct.Get(employee, "Company.Address.Phone"); !found {
		fmt.Println("✓ Correctly handled non-existent field 'Company.Address.Phone'")
	}

	// Try to set non-existent field
	err = rstruct.Set(employeePtr, "Company.Address.Phone", "123-456-7890")
	if err != nil {
		fmt.Printf("✓ Correctly handled setting non-existent field 'Company.Address.Phone': %v\n", err)
	}

	// Try to set wrong type
	err = rstruct.Set(employeePtr, "Age", "thirty")
	if err != nil {
		fmt.Printf("✓ Correctly handled type mismatch for age field: %v\n", err)
	}

	fmt.Println("\n=== Examples Complete ===")
}

func printEmployee(emp Employee) {
	fmt.Printf("  Name: %s\n", emp.Name)
	fmt.Printf("  Age: %d\n", emp.Age)
	fmt.Printf("  Email: %s\n", emp.Email)
	fmt.Printf("  Company: %s\n", emp.Company.Name)
	fmt.Printf("  Company Address: %s, %s, %s\n",
		emp.Company.Address.Street,
		emp.Company.Address.City,
		emp.Company.Address.Country)
	fmt.Printf("  Active: %t\n", emp.Active)
}
