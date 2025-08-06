package main

import (
	"fmt"

	"github.com/jkaveri/ramda"
)

type User struct {
	ID   string
	Name string
	Age  int
}

type Product struct {
	SKU      string
	Name     string
	Price    float64
	Category string
}

func main() {
	fmt.Println("IndexBy Function Examples")
	fmt.Println("========================")

	// Example 1: Index users by ID
	users := []User{
		{ID: "1", Name: "Alice", Age: 25},
		{ID: "2", Name: "Bob", Age: 30},
		{ID: "3", Name: "Charlie", Age: 35},
		{ID: "4", Name: "Diana", Age: 28},
	}

	userIndex := ramda.IndexBy(func(u User) string { return u.ID }, users)
	fmt.Println("\n1. Users indexed by ID:")
	for id, user := range userIndex {
		fmt.Printf("  %s: %s (Age: %d)\n", id, user.Name, user.Age)
	}

	// Example 2: Index products by SKU
	products := []Product{
		{SKU: "PROD-001", Name: "Laptop", Price: 999.99, Category: "Electronics"},
		{SKU: "PROD-002", Name: "Mouse", Price: 29.99, Category: "Electronics"},
		{SKU: "PROD-003", Name: "Desk", Price: 199.99, Category: "Furniture"},
		{SKU: "PROD-004", Name: "Chair", Price: 149.99, Category: "Furniture"},
	}

	productIndex := ramda.IndexBy(func(p Product) string { return p.SKU }, products)
	fmt.Println("\n2. Products indexed by SKU:")
	for sku, product := range productIndex {
		fmt.Printf("  %s: %s ($%.2f) - %s\n", sku, product.Name, product.Price, product.Category)
	}

	// Example 3: Index by category
	categoryIndex := ramda.IndexBy(func(p Product) string { return p.Category }, products)
	fmt.Println("\n3. Products indexed by Category (shows overwrite behavior):")
	for category, product := range categoryIndex {
		fmt.Printf("  %s: %s ($%.2f)\n", category, product.Name, product.Price)
	}

	// Example 4: Index by age group
	ageGroupIndex := ramda.IndexBy(func(u User) string {
		if u.Age < 30 {
			return "Young"
		} else if u.Age < 40 {
			return "Middle-aged"
		}
		return "Senior"
	}, users)
	fmt.Println("\n4. Users indexed by age group:")
	for group, user := range ageGroupIndex {
		fmt.Printf("  %s: %s (Age: %d)\n", group, user.Name, user.Age)
	}

	// Example 5: Index by price range
	priceRangeIndex := ramda.IndexBy(func(p Product) string {
		if p.Price < 50 {
			return "Budget"
		} else if p.Price < 200 {
			return "Mid-range"
		}
		return "Premium"
	}, products)
	fmt.Println("\n5. Products indexed by price range:")
	for priceRange, product := range priceRangeIndex {
		fmt.Printf("  %s: %s ($%.2f)\n", priceRange, product.Name, product.Price)
	}

	// Example 6: Demonstrating overwrite behavior
	duplicateUsers := []User{
		{ID: "1", Name: "Alice", Age: 25},
		{ID: "1", Name: "AliceUpdated", Age: 26}, // Same ID, different name
		{ID: "2", Name: "Bob", Age: 30},
	}

	duplicateIndex := ramda.IndexBy(func(u User) string { return u.ID }, duplicateUsers)
	fmt.Println("\n6. Demonstrating overwrite behavior:")
	for id, user := range duplicateIndex {
		fmt.Printf("  %s: %s (Age: %d)\n", id, user.Name, user.Age)
	}

	// Example 7: Using IndexBy with other functions
	fmt.Println("\n7. Combining IndexBy with other functions:")

	// Filter expensive products, then index by category
	expensiveProducts := ramda.Filter(func(p Product) bool { return p.Price > 100 }, products)
	expensiveByCategory := ramda.IndexBy(func(p Product) string { return p.Category }, expensiveProducts)

	fmt.Println("  Expensive products by category:")
	for category, product := range expensiveByCategory {
		fmt.Printf("    %s: %s ($%.2f)\n", category, product.Name, product.Price)
	}

	// Example 8: Using ToSet for membership checking
	fmt.Println("\n8. Set operations with ToSet:")

	// Create sets from slices
	numbers := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 5}
	numberSet := ramda.ToSet(numbers)

	words := []string{"hello", "world", "hello", "go", "world", "programming", "golang"}
	wordSet := ramda.ToSet(words)

	fmt.Println("  Number set (unique values):")
	for num := range numberSet {
		fmt.Printf("    %d\n", num)
	}

	fmt.Println("  Word set (unique values):")
	for word := range wordSet {
		fmt.Printf("    %s\n", word)
	}

	// Demonstrate membership checking
	testNumbers := []int{1, 3, 5, 7, 9}
	fmt.Println("  Membership checking:")
	for _, num := range testNumbers {
		if _, exists := numberSet[num]; exists {
			fmt.Printf("    %d is in the set\n", num)
		} else {
			fmt.Printf("    %d is NOT in the set\n", num)
		}
	}

	testWords := []string{"hello", "go", "python", "golang"}
	fmt.Println("  Word membership checking:")
	for _, word := range testWords {
		if _, exists := wordSet[word]; exists {
			fmt.Printf("    '%s' is in the set\n", word)
		} else {
			fmt.Printf("    '%s' is NOT in the set\n", word)
		}
	}

	// Example 9: Combining ToSet with other functions
	fmt.Println("\n9. Combining ToSet with other functions:")

	// Filter even numbers, then create a set
	evenNumbers := ramda.Filter(func(n int) bool { return n%2 == 0 }, numbers)
	evenNumberSet := ramda.ToSet(evenNumbers)

	fmt.Println("  Even numbers set:")
	for num := range evenNumberSet {
		fmt.Printf("    %d\n", num)
	}

	// Create a set of product categories
	categorySet := ramda.ToSet(ramda.Map(func(p Product) string { return p.Category }, products))
	fmt.Println("  Product categories set:")
	for category := range categorySet {
		fmt.Printf("    %s\n", category)
	}
}
