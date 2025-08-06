package main

import (
	"fmt"

	"github.com/jkaveri/ramda/rslice"
)

func main() {
	fmt.Println("Slice Operations Examples")
	fmt.Println("=======================")
	fmt.Println()

	// 1. Map - Transform elements
	fmt.Println("1. Map - Transform elements:")
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("  Original: %v\n", numbers)

	doubled := rslice.Map(func(n int) int { return n * 2 }, numbers)
	fmt.Printf("  Doubled: %v\n", doubled)

	strings := []string{"hello", "world", "go"}
	lengths := rslice.Map(func(s string) int { return len(s) }, strings)
	fmt.Printf("  String lengths: %v\n", lengths)
	fmt.Println()

	// 2. Filter - Select elements
	fmt.Println("2. Filter - Select elements:")
	allNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("  Original: %v\n", allNumbers)

	even := rslice.Filter(func(n int) bool { return n%2 == 0 }, allNumbers)
	fmt.Printf("  Even numbers: %v\n", even)

	greaterThan5 := rslice.Filter(func(n int) bool { return n > 5 }, allNumbers)
	fmt.Printf("  Numbers > 5: %v\n", greaterThan5)
	fmt.Println()

	// 3. Reduce - Accumulate values
	fmt.Println("3. Reduce - Accumulate values:")
	fmt.Printf("  Original: %v\n", numbers)

	sum := rslice.Reduce(func(acc, n int) int { return acc + n }, 0, numbers)
	fmt.Printf("  Sum: %d\n", sum)

	product := rslice.Reduce(func(acc, n int) int { return acc * n }, 1, numbers)
	fmt.Printf("  Product: %d\n", product)

	words := []string{"hello", "world", "go", "programming"}
	concatenated := rslice.Reduce(func(acc, s string) string { return acc + " " + s }, "", words)
	fmt.Printf("  Concatenated: %s\n", concatenated)
	fmt.Println()

	// 4. Find - Find first matching element
	fmt.Println("4. Find - Find first matching element:")
	fmt.Printf("  Original: %v\n", allNumbers)

	found, exists := rslice.Find(func(n int) bool { return n > 7 }, allNumbers)
	if exists {
		fmt.Printf("  First number > 7: %d\n", found)
	} else {
		fmt.Println("  No number > 7 found")
	}

	_, exists = rslice.Find(func(n int) bool { return n > 20 }, allNumbers)
	if !exists {
		fmt.Println("  No number > 20 found (as expected)")
	}
	fmt.Println()

	// 5. Any and All - Check conditions
	fmt.Println("5. Any and All - Check conditions:")
	fmt.Printf("  Original: %v\n", allNumbers)

	hasEven := rslice.Any(func(n int) bool { return n%2 == 0 }, allNumbers)
	fmt.Printf("  Has even numbers: %t\n", hasEven)

	allPositive := rslice.All(func(n int) bool { return n > 0 }, allNumbers)
	fmt.Printf("  All positive: %t\n", allPositive)

	allEven := rslice.All(func(n int) bool { return n%2 == 0 }, allNumbers)
	fmt.Printf("  All even: %t\n", allEven)
	fmt.Println()

	// 6. Take and Drop - Slice operations
	fmt.Println("6. Take and Drop - Slice operations:")
	fmt.Printf("  Original: %v\n", allNumbers)

	firstThree := rslice.Take(3, allNumbers)
	fmt.Printf("  First 3: %v\n", firstThree)

	withoutFirstTwo := rslice.Drop(2, allNumbers)
	fmt.Printf("  Without first 2: %v\n", withoutFirstTwo)

	lastThree := rslice.Take(3, rslice.Reverse(allNumbers))
	fmt.Printf("  Last 3: %v\n", lastThree)
	fmt.Println()

	// 7. Unique - Remove duplicates
	fmt.Println("7. Unique - Remove duplicates:")
	duplicates := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}
	fmt.Printf("  Original: %v\n", duplicates)

	unique := rslice.Unique(duplicates)
	fmt.Printf("  Unique: %v\n", unique)

	wordsWithDups := []string{"hello", "world", "hello", "go", "world", "programming"}
	uniqueWords := rslice.Unique(wordsWithDups)
	fmt.Printf("  Unique words: %v\n", uniqueWords)
	fmt.Println()

	// 8. Flatten - Flatten nested slices
	fmt.Println("8. Flatten - Flatten nested slices:")
	matrix := [][]int{{1, 2}, {3, 4}, {5, 6, 7}, {8}}
	fmt.Printf("  Matrix: %v\n", matrix)

	flattened := rslice.Flatten(matrix)
	fmt.Printf("  Flattened: %v\n", flattened)
	fmt.Println()

	// 9. Zip - Combine two slices
	fmt.Println("9. Zip - Combine two slices:")
	numbers2 := []int{1, 2, 3, 4}
	letters := []string{"a", "b", "c", "d", "e"}
	fmt.Printf("  Numbers: %v\n", numbers2)
	fmt.Printf("  Letters: %v\n", letters)

	zipped := rslice.Zip(numbers2, letters)
	fmt.Printf("  Zipped: %v\n", zipped)
	fmt.Println()

	// 10. Reverse - Reverse slice order
	fmt.Println("10. Reverse - Reverse slice order:")
	fmt.Printf("  Original: %v\n", numbers)

	reversed := rslice.Reverse(numbers)
	fmt.Printf("  Reversed: %v\n", reversed)
	fmt.Println()

	// 11. GroupBy - Group elements by key
	fmt.Println("11. GroupBy - Group elements by key:")
	fmt.Printf("  Original: %v\n", allNumbers)

	groupedByParity := rslice.GroupBy(func(n int) string {
		if n%2 == 0 {
			return "even"
		}
		return "odd"
	}, allNumbers)
	fmt.Printf("  Grouped by parity: %v\n", groupedByParity)

	groupedByRange := rslice.GroupBy(func(n int) string {
		switch {
		case n <= 3:
			return "low"
		case n <= 7:
			return "medium"
		default:
			return "high"
		}
	}, allNumbers)
	fmt.Printf("  Grouped by range: %v\n", groupedByRange)
	fmt.Println()

	// 12. SortBy - Sort with custom comparison
	fmt.Println("12. SortBy - Sort with custom comparison:")
	unsorted := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	fmt.Printf("  Original: %v\n", unsorted)

	sorted := rslice.SortBy(func(a, b int) bool { return a < b }, unsorted)
	fmt.Printf("  Sorted ascending: %v\n", sorted)

	sortedDesc := rslice.SortBy(func(a, b int) bool { return a > b }, unsorted)
	fmt.Printf("  Sorted descending: %v\n", sortedDesc)
	fmt.Println()

	// 13. IndexBy - Create map from slice
	fmt.Println("13. IndexBy - Create map from slice:")
	type User struct {
		ID   string
		Name string
		Age  int
	}

	users := []User{
		{ID: "1", Name: "Alice", Age: 25},
		{ID: "2", Name: "Bob", Age: 30},
		{ID: "3", Name: "Charlie", Age: 35},
		{ID: "4", Name: "Diana", Age: 28},
	}

	indexedByID := rslice.IndexBy(func(u User) string { return u.ID }, users)
	fmt.Printf("  Indexed by ID: %v\n", indexedByID)

	indexedByName := rslice.IndexBy(func(u User) string { return u.Name }, users)
	fmt.Printf("  Indexed by Name: %v\n", indexedByName)
	fmt.Println()

	// 14. ToSet - Create set from slice
	fmt.Println("14. ToSet - Create set from slice:")
	numbersWithDups := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 5}
	fmt.Printf("  Original: %v\n", numbersWithDups)

	numberSet := rslice.ToSet(numbersWithDups)
	fmt.Printf("  Set: %v\n", numberSet)

	// Check membership
	if _, exists := numberSet[3]; exists {
		fmt.Println("  3 is in the set")
	}
	if _, exists := numberSet[6]; !exists {
		fmt.Println("  6 is not in the set")
	}
	fmt.Println()

	// 15. Function composition
	fmt.Println("15. Function composition:")
	fmt.Printf("  Original: %v\n", allNumbers)

	// Filter even numbers, double them, take first 3, reverse
	result := rslice.Reverse(
		rslice.Take(3,
			rslice.Map(func(n int) int { return n * 2 },
				rslice.Filter(func(n int) bool { return n%2 == 0 }, allNumbers))))
	fmt.Printf("  Filter even → Double → Take 3 → Reverse: %v\n", result)

	// Find first number > 5, then double it if found
	if found, exists := rslice.Find(func(n int) bool { return n > 5 }, allNumbers); exists {
		doubled := rslice.Map(func(n int) int { return n * 2 }, []int{found})
		fmt.Printf("  First number > 5 doubled: %v\n", doubled)
	}
}
